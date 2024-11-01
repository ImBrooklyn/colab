package handler

import (
    "colab/api/common"
    "colab/api/grpc/snowflake"
    "colab/app/colab-snowflake/internal/global"
    "context"
)

type SnowflakeServer struct {
}

func NewSnowflakeServer() *SnowflakeServer {
    return &SnowflakeServer{}
}

func (s *SnowflakeServer) NextID(_ context.Context, _ *common.Empty) (*snowflake.SnowflakeID, error) {
    id := global.Node.Generate()
    return &snowflake.SnowflakeID{
        Id: id.Int64(),
    }, nil
}
