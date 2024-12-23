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

func GetConfigList(c *gin.Context) {
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
	var req model.GetConfigListReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.UserID = userID
	req.ClientIP = c.ClientIP()
	resp, err := service.GetConfigList(c, req)
	r.JsonReturn(err, resp)
	return
}
func SaveConfig(c *gin.Context) {
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
	var req model.SaveConfigReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.UserID = userID
	req.ClientIP = c.ClientIP()
	resp, err := service.SaveConfig(c, req)
	r.JsonReturn(err, resp)
	return
}

func RefreshConfig(c *gin.Context) {
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
	var req model.RefreshConfigReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.UserID = userID
	req.ClientIP = c.ClientIP()
	resp, err := service.RefreshConfig(c, req)
	r.JsonReturn(err, resp)
	return
}
