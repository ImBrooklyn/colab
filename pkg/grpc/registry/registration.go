package registry

import (
    "colab/pkg/config"
    "fmt"
    consul "github.com/hashicorp/consul/api"
    uuid "github.com/satori/go.uuid"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    hc "google.golang.org/grpc/health/grpc_health_v1"
)

func RegisterService(server *grpc.Server, srvConfig *config.SrvConfig, consulConfig *config.ConsulConfig) (*consul.Client, string) {
    // register health check
    hc.RegisterHealthServer(server, health.NewServer())

    cfg := consul.DefaultConfig()
    cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

    client, err := consul.NewClient(cfg)
    if err != nil {
        panic(err)
    }

    serviceID := fmt.Sprintf("%s", uuid.NewV4())
    registration := &consul.AgentServiceRegistration{
        Name:    srvConfig.AppName,
        ID:      serviceID,
        Port:    srvConfig.Port,
        Tags:    srvConfig.Tags,
        Address: srvConfig.Host,
        Check: &consul.AgentServiceCheck{
            GRPC:                           fmt.Sprintf("%s:%d", srvConfig.Host, srvConfig.Port),
            Timeout:                        "5s",
            Interval:                       "5s",
            DeregisterCriticalServiceAfter: "15s",
        },
    }

    err = client.Agent().ServiceRegister(registration)
    if err != nil {
        panic(err)
    }

    return client, serviceID
}
