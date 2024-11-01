package initialize

import (
    "colab/api/grpc/message"
    "colab/app/colab-message/internal/global"
    "colab/app/colab-message/internal/handler"
    "colab/pkg/grpc/registry"
    consul "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
)

func RegisterService(server *grpc.Server) (*consul.Client, string) {
    message.RegisterMessageServiceServer(server, handler.NewMessageServer())

    return registry.RegisterService(server, global.Config.Srv, global.Config.Consul)
}
