package initialize

import (
    "colab/app/colab-snowflake/conf"
    "colab/app/colab-snowflake/internal/global"
    "colab/app/pkg/constants/keys"
    "colab/pkg/config"
    "github.com/spf13/viper"
    "strconv"
)

func InitConfig() *conf.SnowflakeSrvConfig {
    config.LoadConfig("colab-snowflake", &global.Config)

    viper.AutomaticEnv()
    nodeIdStr := viper.GetString(keys.KSnowflakeNodeId)
    if nodeIdStr == "" {
        nodeIdStr = "1"
    }

    if nodeID, err := strconv.Atoi(nodeIdStr); err != nil {
        global.Config.NodeID = 1
    } else {
        global.Config.NodeID = nodeID
    }

    return global.Config
}
