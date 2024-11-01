package bootstrap

import (
    consul "github.com/hashicorp/consul/api"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "net"
    "os"
    "os/signal"
    "syscall"
)

func RunGRPCServer(consulClient *consul.Client, serviceID string, server *grpc.Server, listener net.Listener) {
    go func() {
        err := server.Serve(listener)
        if err != nil {
            panic("failed to start grpc:" + err.Error())
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    if err := consulClient.Agent().ServiceDeregister(serviceID); err != nil {
        zap.S().Info("deregister failed")
    } else {
        zap.S().Info("deregister success")
    }
}
