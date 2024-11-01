package client

import (
    "colab/pkg/config"
    "fmt"
    _ "github.com/mbobakov/grpc-consul-resolver" // It's important
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func NewSrvClientConn(srvName string, consulConfig *config.ConsulConfig) *grpc.ClientConn {
    opts := []grpc.DialOption{
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
    }

    clientTarget := fmt.Sprintf("consul://%s:%d/%s?wait=14s",
        consulConfig.Host, consulConfig.Port, srvName)
    conn, err := grpc.NewClient(clientTarget, opts...)
    if err != nil {
        panic(fmt.Sprintf("init %s srv failed", srvName))
    }
    return conn
}
