package main

import (
    "colab/app/colab-api/initialize"
    "colab/pkg/logging"
    "fmt"
    "go.uber.org/zap"
)

func main() {
    logging.InitLogger()
    config := initialize.InitConfig()
    router := initialize.InitRouter()
    initialize.InitLocale()
    initialize.InitValidator()
    _ = router
    initialize.InitSrvClient()

    err := router.Run(fmt.Sprintf(":%d", config.Port))
    if err != nil {
        zap.S().Panic("boot failed, err: %s", err.Error())
    }
}
