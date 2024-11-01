package global

import (
    "colab/app/colab-portal/conf"
    "github.com/go-redis/redis/v8"
)

var (
    Config *conf.PortalSrvConfig
    Redis  *redis.Client
)
