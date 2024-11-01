package conf

import "colab/pkg/config"

type SnowflakeSrvConfig struct {
    Consul *config.ConsulConfig `mapstructure:"consul" yaml:"consul" json:"consul"`
    Srv    *config.SrvConfig    `mapstructure:"srv" yaml:"srv" json:"srv"`
    NodeID int                  `mapstructure:"node-id" yaml:"node-id" json:"node_id"`
}
