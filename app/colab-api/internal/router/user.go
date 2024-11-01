package router

import (
    "colab/app/colab-api/internal/handler"
    "github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
    group := router.Group("/user")

    {
        group.POST("/signup", handler.Signup)
        group.POST("/signin", handler.Signin)
    }
}

func InitOrgRouter(router *gin.RouterGroup) {
    group := router.Group("/organization")

    {
        group.GET("", handler.GetOrganizationList)
    }
}
