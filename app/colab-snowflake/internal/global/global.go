package global

import (
    "colab/app/colab-snowflake/conf"
    "github.com/bwmarrin/snowflake"
)

var (
    Config *conf.SnowflakeSrvConfig
    Node   *snowflake.Node
)
