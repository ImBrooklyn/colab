package initialize

import (
    "colab/app/colab-message/conf"
    "colab/app/colab-message/internal/global"
    "colab/pkg/config"
)

func InitConfig() *conf.MessageSrvConfig {
    config.LoadConfig("colab-message", &global.Config)
    return global.Config
}
