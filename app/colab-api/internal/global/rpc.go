package global

import (
    "colab/api/grpc/message"
    "colab/api/grpc/portal"
    "colab/api/grpc/user"
)

var (
    MessageSrvClient      message.MessageServiceClient
    UserSrvClient         user.UserServiceClient
    UserSignSrvClient     user.UserSignServiceClient
    OrganizationSrvClient user.OrganizationServiceClient

    MenuSrvClient portal.MenuServiceClient
)
