package handler

import (
    "colab/api/grpc/user"
    "colab/app/colab-api/internal/global"
    "colab/app/colab-api/internal/util"
    "colab/app/pkg/constants/keys"
    "colab/app/pkg/model/rest"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetOrganizationList(c *gin.Context) {
    ctx, cancel := util.NewApiContext(c)
    defer cancel()
    value := ctx.Value(keys.KCtxUseID)
    fmt.Println(value)

    request := &user.OrganizationListRequest{}
    response, err := global.OrganizationSrvClient.GetOrganizationList(ctx, request)
    if err != nil {
        util.HandleGrpcError(c, err)
        return
    }

    c.JSON(http.StatusOK, rest.Success(response))
}
