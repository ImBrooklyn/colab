package main

import (
    "colab/app/colab-message/initialize"
    "colab/pkg/bootstrap"
    "colab/pkg/grpc/network"
    "colab/pkg/logging"
    "google.golang.org/grpc"
)

func main() {
    logging.InitLogger()
    config := initialize.InitConfig()
    initialize.InitMiddlewares(config)

    listener := network.InitNetwork(config.Srv.Port)
    server := grpc.NewServer()
    consulClient, serviceID := initialize.RegisterService(server)

    bootstrap.RunGRPCServer(consulClient, serviceID, server, listener)
}
