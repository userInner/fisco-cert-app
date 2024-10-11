package redis

import (
	"fisco-cert-app/setting"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

// Init 初始化redis
func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result() // 检测redis服务是否存活
	if err != nil {
		return
	}
	zap.L().Info("redis connected success!")
	return
}

// Close 手动关闭close
func Close() {
	_ = client.Close()
}
