package e

import "TelegramBot/pkg/xlerror"

var (
	ErrRequest = xlerror.ErrRequest
	ErrServer  = xlerror.ErrServer
	ErrForWait = xlerror.ErrForWait
	ErrorToken = xlerror.ErrToken

	ErrorUsernameOrPassword = xlerror.New(10001, "用户名或密码错误")
)
