package initialize

import (
    "colab/api/grpc/snowflake"
    "colab/api/grpc/user"
    "colab/app/colab-user/internal/global"
    "colab/app/colab-user/internal/handler"
    "colab/app/pkg/constants/srv"
    "colab/pkg/grpc/client"
    "colab/pkg/grpc/registry"
    consul "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
)

func RegisterService(server *grpc.Server) (*consul.Client, string) {
    user.RegisterUserSignServiceServer(server, handler.NewUserSignServer())
    user.RegisterUserServiceServer(server, handler.NewUserServer())
    user.RegisterOrganizationServiceServer(server, handler.NewOrganizationServer())

    return registry.RegisterService(server, global.Config.Srv, global.Config.Consul)
}

func initSnowflakeSrvConn() {
    conn := client.NewSrvClientConn(srv.SnowflakeSrv, global.Config.Consul)
    global.Snowflake = snowflake.NewSnowflakeServiceClient(conn)
}

func InitSrvClient() {
    initSnowflakeSrvConn()
}
