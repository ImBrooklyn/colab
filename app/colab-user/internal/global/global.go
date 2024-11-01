package global

import (
    "colab/app/colab-user/conf"
    "github.com/go-redis/redis/v8"
)

var (
    Config *conf.UserSrvConfig
    Redis  *redis.Client
)
