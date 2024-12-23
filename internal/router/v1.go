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
	// web hook
	hookGroup := apiGroup.Group("hook")
	{
		hookGroup.POST("telegram-bot/:name", v1.TelegramBotWebHook)
	}
	loginGroup := apiGroup.Group("login")
	{
		loginGroup.POST("login", v1.AdminLogin)
	}
	apiGroup.Use(xlhttp.JWTHeadMiddleware(xlhttp.NewJWT(config.C.Web.Secret, time.Second*3600)))
	apiGroup.POST("/login/change_password", v1.ChangePassword) // 修改密码
	botGroup := apiGroup.Group("bot")
	{
		// 所有机器人
		botGroup.POST("all", v1.GetAllBots)
		// 机器人列表
		botGroup.POST("list", v1.GetBotList)
		// 保存机器人信息
		botGroup.POST("save", v1.SaveBot)
	}
	actionGroup := apiGroup.Group("action")
	{
		// action列表
		actionGroup.POST("list", v1.GetActionList)
		// 保存action
		actionGroup.POST("save", v1.SaveAction)
		// 删除action
		actionGroup.POST("delete", v1.DeleteAction)
	}
	configGroup := apiGroup.Group("config")
	{
		// config列表
		configGroup.POST("list", v1.GetConfigList)
		// 保存config
		configGroup.POST("save", v1.SaveConfig)
		// 刷新config
		configGroup.POST("refresh", v1.RefreshConfig)
	}
	groupGroup := apiGroup.Group("group")
	{
		// group列表
		groupGroup.POST("list", v1.GetGroupList)
		// 保存group
		groupGroup.POST("save", v1.SaveGroup)
		// 删除group
		groupGroup.POST("delete", v1.DeleteGroup)
	}
	messageGroup := apiGroup.Group("message")
	{
		// message列表
		messageGroup.POST("list", v1.GetMessagePushList)
		// 保存message
		messageGroup.POST("save", v1.SaveMessagePush)
	}
}
