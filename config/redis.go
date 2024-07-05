package config

import (
	"TelegramBot/internal/dao"
	"TelegramBot/pkg/xlredis"
)

func initRedis(c RedisConfig) {
	r, err := xlredis.NewClient(c.Hosts, c.Username, c.Password, c.Prefix, c.DB)
	if err != nil {
		panic(err)
	}
	dao.Redis = r
}
