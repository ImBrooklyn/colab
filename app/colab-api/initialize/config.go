package initialize

import (
    "colab/app/colab-api/conf"
    "colab/app/colab-api/internal/global"
    "colab/pkg/config"
)

func InitConfig() *conf.ApiServerConfig {
    config.LoadConfig("colab-api", &global.Config)
    return global.Config
}
