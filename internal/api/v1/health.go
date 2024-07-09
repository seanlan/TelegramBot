//generated by lazy
//author: seanlan

package v1

import (
	"TelegramBot/internal/model"
	"TelegramBot/internal/service"
	"TelegramBot/pkg/xlhttp"
	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	var (
		err error
	)
	r := xlhttp.Build(c)
	var req model.CheckHealthReq
	err = r.RequestParser(&req)
	if err != nil {
		return
	}
	req.ClientIP = c.ClientIP()
	resp, err := service.CheckHealth(c, req)
	r.JsonReturn(err, resp)
	return
}
