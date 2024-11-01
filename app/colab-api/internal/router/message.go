package router

import (
    "colab/app/colab-api/internal/handler"
    "github.com/gin-gonic/gin"
)

func InitMessageRouter(router *gin.RouterGroup) {
    group := router.Group("/message")
    {
        group.POST("/sms-code", handler.SendSmsCode)
    }
}
