package initialize

import (
    "colab/app/colab-portal/conf"
    "colab/app/colab-portal/internal/global"
    "colab/pkg/config"
)

func InitConfig() *conf.PortalSrvConfig {
    config.LoadConfig("colab-portal", &global.Config)
    if global.Config == nil {
        panic("Init config failed")
    }
    return global.Config
}
