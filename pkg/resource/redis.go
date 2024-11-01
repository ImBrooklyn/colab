package resource

import (
    "colab/pkg/config"
    "github.com/go-redis/redis/v8"
)

func NewRedisConn(cfg *config.RedisConfig) *redis.Client {
    opt := &redis.Options{
        Addr: cfg.Addr(),
    }
    return redis.NewClient(opt)
}
