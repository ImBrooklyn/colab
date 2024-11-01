package conf

import (
    "colab/pkg/config"
)

type UserSrvConfig struct {
    Consul *config.ConsulConfig `mapstructure:"consul" yaml:"consul" json:"consul"`
    Srv    *config.SrvConfig    `mapstructure:"srv" yaml:"srv" json:"srv"`
    Redis  *config.RedisConfig  `mapstructure:"redis" yaml:"redis" json:"redis"`
    Mysql  *config.MysqlConfig  `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
}
