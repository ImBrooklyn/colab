package initialize

import (
    "colab/app/colab-snowflake/conf"
    "colab/app/colab-snowflake/internal/global"
    "github.com/bwmarrin/snowflake"
    "go.uber.org/zap"
    "time"
)

func InitNode(config *conf.SnowflakeSrvConfig) {
    st, _ := time.Parse("2006-01-02", "2024-01-01")
    snowflake.Epoch = st.UnixMilli()
    zap.S().Infof("snowflake epoch set: %d", snowflake.Epoch)
    node, err := snowflake.NewNode(int64(config.NodeID))
    if err != nil {
        panic(err)
    }

    global.Node = node
    generate := node.Generate()
    generate.Int64()
}
