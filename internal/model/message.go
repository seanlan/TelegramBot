//generated by lazy
//author: seanlan

package model

import "TelegramBot/internal/dao/sqlmodel"

type GetMessagePushListReq struct {
	BaseReq
	BotID int64 `json:"bot_id" form:"bot_id"`
	Page  int   `json:"page" form:"page" binding:"required"`
	Size  int   `json:"size" form:"size" binding:"required"`
}

type GetMessagePushListResp struct {
	List  []sqlmodel.BotMessagePush `json:"list"`
	Total int64                     `json:"total"`
}

type SaveMessagePushReq struct {
	BaseReq
	ID        int64  `json:"id" form:"id"`
	BotID     int64  `json:"bot_id" form:"bot_id" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	Image     string `json:"image" form:"image" binding:"omitempty"`
	Extension string `json:"extension" form:"extension"`
	StartAt   int64  `json:"start_at" form:"start_at" binding:"required"`
}

type SaveMessagePushResp struct {
}