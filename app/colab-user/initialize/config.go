package initialize

import (
    "colab/app/colab-user/conf"
    "colab/app/colab-user/internal/global"
    "colab/pkg/config"
)

func InitConfig() *conf.UserSrvConfig {
    config.LoadConfig("colab-user", &global.Config)
    if global.Config == nil {
        panic("Init config failed")
    }
    return global.Config
}
