package config

import "fmt"

type NacosConfig struct {
    Host      string `mapstructure:"host"`
    Port      uint64 `mapstructure:"port"`
    Namespace string `mapstructure:"namespace"`
    Username  string `mapstructure:"username"`
    Password  string `mapstructure:"password"`
    DataId    string `mapstructure:"data-id"`
    Group     string `mapstructure:"group"`
}

type SrvConfig struct {
    AppName string   `mapstructure:"app-name" yaml:"app-name" json:"app_name"`
    Host    string   `mapstructure:"host" yaml:"host" json:"host"`
    Port    int      `mapstructure:"port" yaml:"port" json:"port"`
    Tags    []string `mapstructure:"tags" yaml:"tags" json:"tags"`
}

type JWTConfig struct {
    AccessSecret string `mapstructure:"access-secret" yaml:"access-secret" json:"access_secret"`
    Expiry       int    `mapstructure:"expiry" yaml:"expiry" json:"expiry"`
}

type ConsulConfig struct {
    Host string `mapstructure:"host" yaml:"host" json:"host"`
    Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

type RedisConfig struct {
    Host string `mapstructure:"host" yaml:"host" json:"host"`
    Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

func (rc *RedisConfig) Addr() string {
    return fmt.Sprintf("%s:%d", rc.Host, rc.Port)
}

type MysqlConfig struct {
    Host     string `mapstructure:"host" yaml:"host" json:"host"`
    Port     int    `mapstructure:"port" yaml:"port" json:"port"`
    Database string `mapstructure:"database" yaml:"database" json:"database"`
    Username string `mapstructure:"username" yaml:"username" json:"username"`
    Password string `mapstructure:"password" yaml:"password" json:"password"`
}
