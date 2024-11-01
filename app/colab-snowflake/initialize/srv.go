package initialize

import (
    "colab/api/grpc/snowflake"
    "colab/app/colab-snowflake/internal/global"
    "colab/app/colab-snowflake/internal/handler"
    "colab/pkg/grpc/registry"
    consul "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
)

func RegisterService(server *grpc.Server) (*consul.Client, string) {
    snowflake.RegisterSnowflakeServiceServer(server, handler.NewSnowflakeServer())

    return registry.RegisterService(server, global.Config.Srv, global.Config.Consul)
}
