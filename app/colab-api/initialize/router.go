package initialize

import (
    "colab/app/colab-api/internal/middleware"
    "colab/app/colab-api/internal/router"
    "github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    g := gin.Default()

    g.Use(middleware.Cors())
    g.Use(middleware.JWTAuth())

    apiGrp := g.Group("/colab")
    router.InitUserRouter(apiGrp)
    router.InitOrgRouter(apiGrp)
    router.InitMessageRouter(apiGrp)
    router.InitPortalRouter(apiGrp)
    return g
}
