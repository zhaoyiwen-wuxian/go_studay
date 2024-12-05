package cache

import (
	"context"
	"fmt"
	"go_dome/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	cfg := config.AppConfig.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
}
