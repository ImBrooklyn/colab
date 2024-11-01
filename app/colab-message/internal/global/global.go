package global

import (
    "colab/app/colab-message/conf"
    "github.com/go-redis/redis/v8"
)

var (
    Config *conf.MessageSrvConfig
    Redis  *redis.Client
)
