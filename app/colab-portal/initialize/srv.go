package initialize

import (
    "colab/api/grpc/portal"
    "colab/app/colab-portal/internal/global"
    "colab/app/colab-portal/internal/handler"
    "colab/pkg/grpc/registry"
    consul "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
)

func RegisterService(server *grpc.Server) (*consul.Client, string) {
    portal.RegisterMenuServiceServer(server, handler.NewMenuServer())

    return registry.RegisterService(server, global.Config.Srv, global.Config.Consul)
}
