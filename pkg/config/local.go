package config

import (
    "fmt"
    "github.com/spf13/viper"
)

func loadLocalConfig(appName string, env string, cfg any) {
    configFileName := fmt.Sprintf("configs/local/%s-%s.yaml", appName, env)
    v := viper.New()
    v.SetConfigFile(configFileName)
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }

    if err := v.Unmarshal(cfg); err != nil {
        panic(err)
    }
}
