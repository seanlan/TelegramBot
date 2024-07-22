package crontab

import (
	"TelegramBot/internal/dao"
	"TelegramBot/internal/dao/sqlmodel"
	"TelegramBot/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"sync"
	"time"
)

const (
	sendLimit        = 30
	annLockKey       = "cron:push_message:lock"
	anLastGroupIDKey = "cron:push_message:last_group_id"
)

func SendTelegramMessage(ctx context.Context, message *sqlmodel.BotMessagePush, bot *sqlmodel.Bots, tgGroupID int64) {
	var (
		ext      = make([]model.MessageExt, 0)
		rows     = make([]tele.Row, 0)
		selector = &tele.ReplyMarkup{}
	)
	if message.Extension != "" {
		_err := json.Unmarshal([]byte(message.Extension), &ext)
		if _err != nil {
			zap.S().Infof("unmarshal extension error: %v", _err)
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
	pref := tele.Settings{
		Token:  bot.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		zap.S().Warnf("new bot error: %v", err)
		return
	}
	to := &tele.User{ID: tgGroupID}
	if message.Image != "" {
		_, err = b.Send(to,
			&tele.Photo{
				File:    tele.FromURL(message.Image),
				Caption: message.Content,
			}, selector)
	} else {
		_, err = b.Send(to, message.Content, selector)
	}
	if err != nil {
		zap.S().Warnf("send telegram message error: %v", err)
	} else {
		zap.S().Infof("send telegram message success: %d", tgGroupID)
	}
}

func PushAnnouncement(ctx context.Context, message *sqlmodel.BotMessagePush) (err error) {
	lockKey := fmt.Sprintf("%s", annLockKey)
	if !dao.Redis.GetLock(ctx, lockKey, time.Minute) {
		return
	}
	defer func() {
		dao.Redis.ReleaseLock(ctx, lockKey)
	}()
	var (
		botQ           = sqlmodel.BotsColumns
		groupQ         = sqlmodel.BotGroupsColumns
		bot            sqlmodel.Bots
		groups         []sqlmodel.BotGroups
		lastGroupIDKey = fmt.Sprintf("%s:%d", anLastGroupIDKey, message.ID)
	)
	err = dao.FetchBots(ctx, &bot, dao.And(botQ.ID.Eq(message.BotID)))
	if err != nil {
		return
	}
	lastGroupID, err := dao.Redis.Get(ctx, lastGroupIDKey).Int64()
	if err != nil {
		lastGroupID = 0
	}
	err = dao.FetchAllBotGroups(ctx, &groups,
		dao.And(
			groupQ.ID.Gt(lastGroupID)),
		0, sendLimit, groupQ.ID.Asc())
	if err != nil {
		return
	}
	if len(groups) == 0 {
		message.State = 2
		_, _ = dao.UpdateBotMessagePush(ctx, message)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(groups))
	for i := range groups {
		go func(i int) {
			defer wg.Done()
			zap.S().Infof("push message: %d ", groups[i].ID)
			SendTelegramMessage(ctx, message, &bot, groups[i].GroupTgID)
		}(i)
	}
	wg.Wait()
	lastGroupID = groups[len(groups)-1].ID
	dao.Redis.Set(ctx, lastGroupIDKey, lastGroupID, time.Hour) // 记录最后一条提现记录ID
	return
}

func CronPushGroupMessage() {
	var (
		ctx      = context.TODO()
		messageQ = sqlmodel.BotMessagePushColumns
		message  sqlmodel.BotMessagePush
		err      error
	)
	// 查询公告信息
	err = dao.FetchBotMessagePush(ctx, &message,
		dao.And(
			messageQ.StartAt.Lte(time.Now().Unix()*1000),
			messageQ.State.Neq(2),
		))
	if err != nil {
		zap.S().Infof("fetch MessagePush error: %v", err)
		return
	}
	if message.ID == 0 {
		return
	}
	if message.State == 0 { // 未推送 0
		message.State = 1
		_, _ = dao.UpdateBotMessagePush(ctx, &message)
	}
	// 推送公告
	err = PushAnnouncement(ctx, &message)
	if err != nil {
		zap.S().Warnf("push announcement error: %v", err)
	}
	return
}
