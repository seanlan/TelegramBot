//generated by lazy
//author: seanlan

package model

import "TelegramBot/internal/dao/sqlmodel"

type GetBotListReq struct {
	BaseReq
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type GetBotListResp struct {
	Total int64           `json:"total"`
	Bots  []sqlmodel.Bots `json:"bots"`
}

type SaveBotReq struct {
	BaseReq
	ID     int64  `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Token  string `json:"token" form:"token"`
	Enable int32  `json:"enable" form:"enable"`
}

type SaveBotResp struct {
}
