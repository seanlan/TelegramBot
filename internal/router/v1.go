package router

import (
	"TelegramBot/config"
	v1 "TelegramBot/internal/api/v1"
	"TelegramBot/pkg/gin_zap"
	"TelegramBot/pkg/xlhttp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func initWebRouter(r *gin.Engine) {
	r.Use(
		gin_zap.Ginzap(zap.L(), time.RFC3339, false),
		gin_zap.RecoveryWithZap(zap.L(), true),
	)
	// 接口处理
	apiGroup := r.Group("api/v1")
	// 监控监测
	apiGroup.GET("health", v1.CheckHealth)
	loginGroup := apiGroup.Group("login")
	{
		loginGroup.POST("login", v1.AdminLogin)
	}
	// web hook
	hookGroup := apiGroup.Group("hook")
	{
		hookGroup.POST("telegram-bot/:name", v1.TelegramBotWebHook)
	}
	apiGroup.Use(xlhttp.JWTHeadMiddleware(xlhttp.NewJWT(config.C.Web.Secret, time.Second*3600)))
	apiGroup.POST("/login/change_password", v1.ChangePassword) // 修改密码
}
