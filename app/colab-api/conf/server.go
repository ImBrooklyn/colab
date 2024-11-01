package conf

import (
    "colab/pkg/config"
)

type ApiServerConfig struct {
    AppName string               `mapstructure:"app-name" yaml:"app-name" json:"app_name"`
    Port    int                  `mapstructure:"port" yaml:"port" json:"port"`
    Locale  string               `mapstructure:"locale" yaml:"locale" json:"locale"`
    JWT     *config.JWTConfig    `mapstructure:"jwt" yaml:"jwt" json:"jwt"`
    Consul  *config.ConsulConfig `mapstructure:"consul" yaml:"consul" json:"consul"`
}
