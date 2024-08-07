package handler

import (
	"TelegramBot/internal/constant"
	"TelegramBot/internal/dao"
	"TelegramBot/internal/dao/sqlmodel"
	"TelegramBot/internal/model"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"time"
)

var (
	BotActionsKey = "bot-actions:"
)

// GetBotActions
//
//	@Description: 获取机器人的所有动作
//	@param ctx
//	@param name
//	@return actions
//	@return err
func GetBotActions(ctx context.Context, name string) (actions []sqlmodel.Action, err error) {
	actionKey := BotActionsKey + name
	actionCache := dao.Redis.Get(ctx, actionKey).Val()
	if actionCache == "" {
		var (
			actionQ = sqlmodel.ActionColumns
		)
		err = dao.FetchAllAction(ctx, &actions, actionQ.BotName.Eq(name), 0, 0)
		if err != nil {
			return
		}
		var actionsBytes []byte
		actionsBytes, err = json.Marshal(actions)
		dao.Redis.Set(ctx, actionKey, string(actionsBytes), time.Duration(0))
	} else {
		err = json.Unmarshal([]byte(actionCache), &actions)
	}
	return
}

// ClearBotActionsCache
//
//	@Description: 清理机器人动作缓存
//	@param ctx
//	@param name
func ClearBotActionsCache(ctx context.Context, name string) {
	actionKey := BotActionsKey + name
	dao.Redis.Del(ctx, actionKey)
}

// RegisterBotWebhook
//
//	@Description: 注册机器人的webhook
//	@param ctx
func RegisterBotWebhook(ctx context.Context) {
	var (
		bots []sqlmodel.Bots
	)
	err := dao.FetchAllBots(ctx, &bots, nil, 0, 0)
	if err != nil {
		return
	}
	for _, bot := range bots {
		if bot.Enable == 0 { // 是否启用
			continue
		}
		b, _ := tele.NewBot(tele.Settings{
			Token:   bot.Token,
			Offline: true,
		})
		if b != nil {
			hook, _ := GetConfigByKey(ctx, constant.TelegramBotHookKey)
			zap.S().Infof("hook: %s", hook+bot.Name)
			if hook != "" {
				_err := b.SetWebhook(&tele.Webhook{
					Listen:   ":80",
					Endpoint: &tele.WebhookEndpoint{PublicURL: hook + bot.Name},
				})
				zap.S().Infof("set webhook error: %v", _err)
			}
		}
	}
}

// BotProcessUpdate
//
//	@Description: 处理机器人消息
//	@param ctx
//	@param name
//	@param b
//	@param u
func BotProcessUpdate(ctx context.Context, name string, b *tele.Bot, u *tele.Update) {
	actions, err := GetBotActions(ctx, name)
	if err != nil || len(actions) == 0 {
		zap.S().Warnf("get bot actions error: %v", err)
		return
	}
	for _, action := range actions {
		zap.S().Infof("action: %v", action)
		var (
			ext      = make([]model.MessageExt, 0)
			rows     = make([]tele.Row, 0)
			selector = &tele.ReplyMarkup{}
			image    = action.Image
			content  = action.Content
		)
		if action.Extension != "" {
			_err := json.Unmarshal([]byte(action.Extension), &ext)
			if _err != nil {
				zap.S().Infof("unmarshal extension error: %v", _err)
				continue
			}
			zap.S().Infof("extension: %v", ext)
			for _, e := range ext {
				if e.Type == model.MessageTypeUrl {
					rows = append(rows, selector.Row(selector.URL(e.Text, e.Url)))
				} else if e.Type == model.MessageTypeWebapp {
					rows = append(rows, selector.Row(selector.WebApp(e.Text, &tele.WebApp{URL: e.Url})))
				}
			}
			selector.Inline(rows...)
		}
		b.Handle(action.Command, func(c tele.Context) error {
			if image != "" {
				return c.Send(&tele.Photo{
					File:    tele.FromURL(image),
					Caption: content,
				}, selector)
			} else {
				return c.Send(content, selector)
			}
		})
	}
	b.ProcessUpdate(*u)
	return
}
