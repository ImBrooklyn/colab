package config

import (
    "colab/app/pkg/constants/keys"
    "github.com/spf13/viper"
    "go.uber.org/zap"
)

func LoadConfig(appName string, cfg any) {
    viper.AutomaticEnv()

    env := viper.GetString(keys.KEnvVariable)
    if env == "" {
        env = "dev"
    }
    zap.S().Infof("%s: %s", keys.KEnvVariable, env)

    src := viper.GetString(keys.KCfgSrc)
    if src == "" {
        src = "local"
    }

    switch src {
    case "nacos":
        loadNacosConfig(appName, env, cfg)
    case "local":
        loadLocalConfig(appName, env, cfg)
    default:
        loadLocalConfig(appName, env, cfg)
    }

}
