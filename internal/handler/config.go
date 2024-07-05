package handler

import (
	"TelegramBot/internal/dao"
	"TelegramBot/internal/dao/sqlmodel"
	"context"
	"encoding/json"
	"time"
)

const (
	ConfigCacheKey     = "config"
	TelegramBotHookKey = "telegram-bot-webhook"
)

// GetConfig
//
//	@Description: 获取配置
//	@param ctx
//	@return m
//	@return err
func GetConfig(ctx context.Context) (m map[string]string, err error) {
	configCache := dao.Redis.Get(ctx, ConfigCacheKey).Val()
	if configCache == "" {
		var records []sqlmodel.Config
		err = dao.FetchAllConfig(ctx, &records, nil, 0, 0)
		m = make(map[string]string)
		for _, record := range records {
			m[record.Key] = record.Value
		}
		var configBytes []byte
		configBytes, err = json.Marshal(m)
		dao.Redis.Set(ctx, ConfigCacheKey, string(configBytes), time.Hour*1)
	} else {
		err = json.Unmarshal([]byte(configCache), &m)
	}
	return
}

// GetConfigByKey
//
//	@Description: 通过key获取配置
//	@param ctx
//	@param key
//	@return value
//	@return err
func GetConfigByKey(ctx context.Context, key string) (value string, err error) {
	config, err := GetConfig(ctx)
	if err != nil {
		return
	}
	value = config[key]
	return
}

// ClearConfigCache
//
//	@Description: 清除配置缓存
//	@param ctx
func ClearConfigCache(ctx context.Context) {
	dao.Redis.Del(ctx, ConfigCacheKey)
}
