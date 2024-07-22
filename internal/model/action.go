//generated by lazy
//author: seanlan

package model

import "TelegramBot/internal/dao/sqlmodel"

type GetActionListReq struct {
	BaseReq
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type GetActionListResp struct {
	Total   int64             `json:"total"`
	Actions []sqlmodel.Action `json:"actions"`
}

type SaveActionReq struct {
	BaseReq
	ID        int64  `json:"id" form:"id"`
	BotName   string `json:"bot_name" form:"bot_name"`
	Command   string `json:"command" form:"command"`
	Image     string `json:"image" form:"image"`
	Content   string `json:"content" form:"content"`
	Extension string `json:"extension" form:"extension"`
}

type SaveActionResp struct {
}

type DeleteActionReq struct {
	BaseReq
	ID int64 `json:"id" form:"id" binding:"required"`
}

type DeleteActionResp struct {
}
