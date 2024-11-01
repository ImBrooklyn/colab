package initialize

import (
    "colab/api/grpc/message"
    "colab/api/grpc/portal"
    "colab/api/grpc/user"
    "colab/app/colab-api/internal/global"
    "colab/app/pkg/constants/srv"
    "colab/pkg/grpc/client"
)

func initMessageSrvConn() {
    conn := client.NewSrvClientConn(srv.MessageSrv, global.Config.Consul)

    global.MessageSrvClient = message.NewMessageServiceClient(conn)
}

func initUserSrvConn() {
    conn := client.NewSrvClientConn(srv.UserSrv, global.Config.Consul)

    global.UserSrvClient = user.NewUserServiceClient(conn)
    global.UserSignSrvClient = user.NewUserSignServiceClient(conn)
    global.OrganizationSrvClient = user.NewOrganizationServiceClient(conn)
}

func initPortalSrvConn() {
    conn := client.NewSrvClientConn(srv.PortalSrv, global.Config.Consul)

    global.MenuSrvClient = portal.NewMenuServiceClient(conn)
}

func InitSrvClient() {
    initMessageSrvConn()
    initUserSrvConn()
    initPortalSrvConn()
}
