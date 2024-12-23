//generated by lazy
//author: seanlan

package v1

import (
	"TelegramBot/internal/e"
	"TelegramBot/internal/model"
	"TelegramBot/internal/service"
	"TelegramBot/pkg/xlhttp"
	"github.com/gin-gonic/gin"
)

func GetMessagePushList(c *gin.Context) {
	var (
		err    error
		userID int64
	)
	r := xlhttp.Build(c)
	userID, err = r.GetJWTUID()
	if userID == 0 || err != nil {
		r.JsonReturn(e.ErrorToken)
		return
	}
	var req model.GetMessagePushListReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.UserID = userID
	req.ClientIP = c.ClientIP()
	resp, err := service.GetMessagePushList(c, req)
	r.JsonReturn(err, resp)
	return
}
func SaveMessagePush(c *gin.Context) {
	var (
		err    error
		userID int64
	)
	r := xlhttp.Build(c)
	userID, err = r.GetJWTUID()
	if userID == 0 || err != nil {
		r.JsonReturn(e.ErrorToken)
		return
	}
	var req model.SaveMessagePushReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.UserID = userID
	req.ClientIP = c.ClientIP()
	resp, err := service.SaveMessagePush(c, req)
	r.JsonReturn(err, resp)
	return
}
