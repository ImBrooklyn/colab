package handler

import (
    "colab/api/grpc/portal"
    "colab/app/colab-api/internal/global"
    "colab/app/colab-api/internal/util"
    "colab/app/pkg/model/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

func PortalMenu(c *gin.Context) {
    ctx, cancel := util.NewApiContext(c)
    defer cancel()
    response, err := global.MenuSrvClient.GetUserMenus(ctx, &portal.UserMenuRequest{})
    if err != nil {
        util.HandleGrpcError(c, err)
        return
    }

    c.JSON(http.StatusOK, rest.Success(response.Menus))
}
