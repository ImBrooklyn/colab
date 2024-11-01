package handler

import (
    "colab/api/grpc/message"
    "colab/app/colab-api/internal/global"
    "colab/app/colab-api/internal/model/form"
    "colab/app/colab-api/internal/util"
    "colab/app/pkg/model/rest"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SendSmsCode(c *gin.Context) {
    sendSmsForm := form.SendSmsForm{}
    if err := c.ShouldBind(&sendSmsForm); err != nil {
        util.HandleValidatorError(c, err)
        return
    }

    ctx, cancel := util.NewApiContext(c)
    defer cancel()
    response, err := global.MessageSrvClient.SendSmsCode(ctx, &message.SendSmsCodeRequest{
        Mobile: sendSmsForm.Mobile,
        Biz:    sendSmsForm.Biz,
    })
    if err != nil {
        util.HandleGrpcError(c, err)
        return
    }

    c.JSON(http.StatusOK, rest.Success(response))
}
