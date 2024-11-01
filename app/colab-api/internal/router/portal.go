package router

import (
    "colab/app/colab-api/internal/handler"
    "github.com/gin-gonic/gin"
)

func InitPortalRouter(router *gin.RouterGroup) {
    group := router.Group("/portal")

    {
        group.GET("/menu", handler.PortalMenu)
    }
}
