package initialize

import (
	"context"
	"fmt"
	"gomap/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.EWA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // 不设置密码
		DB:       redisCfg.DB,       // 使用默认的
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.EWA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("====5-redis====: redis init success")
		global.EWA_LOG.Info("redis connect ping response", zap.String("pong", pong))
		global.EWA_REDIS = client
	}
}
