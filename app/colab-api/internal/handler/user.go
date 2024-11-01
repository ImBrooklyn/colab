package handler

import (
    "colab/api/grpc/user"
    "colab/app/colab-api/internal/constants"
    "colab/app/colab-api/internal/global"
    "colab/app/colab-api/internal/middleware"
    "colab/app/colab-api/internal/model/form"
    "colab/app/colab-api/internal/model/resp"
    "colab/app/colab-api/internal/util"
    "colab/app/pkg/errs"
    "colab/app/pkg/model/rest"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net/http"
)

func Signup(c *gin.Context) {
    signupForm := form.SignupForm{}
    if err := c.ShouldBind(&signupForm); err != nil {
        util.HandleValidatorError(c, err)
        return
    }

    ctx, cancel := util.NewApiContext(c)
    defer cancel()

    request := user.SignupRequest{
        Email:    signupForm.Email,
        Username: signupForm.Username,
        Nickname: signupForm.Nickname,
        Password: signupForm.Password,
        Mobile:   signupForm.Mobile,
        SmsCode:  signupForm.SmsCode,
    }
    if _, err := global.UserSignSrvClient.Signup(ctx, &request); err != nil {
        util.HandleGrpcError(c, err)
        return
    }

    c.JSON(http.StatusOK, rest.Success(nil))
}

func Signin(c *gin.Context) {
    signinForm := form.SigninForm{}
    if err := c.ShouldBind(&signinForm); err != nil {
        util.HandleValidatorError(c, err)
        return
    }

    ctx, cancel := util.NewApiContext(c)
    defer cancel()

    var response *user.SigninResponse
    var err error
    switch signinForm.SigninType {
    case constants.SignByUsername, constants.SignByEmail:
        request := &user.SigninByPasswordRequest{
            Account:    signinForm.Identifier,
            Password:   signinForm.Crypto,
            SigninType: signinForm.SigninType,
        }
        response, err = global.UserSignSrvClient.SigninByPassword(ctx, request)
    case constants.SignByMobile:
        request := &user.SigninByMobileRequest{
            Mobile:  signinForm.Identifier,
            SmsCode: signinForm.Crypto,
        }
        response, err = global.UserSignSrvClient.SigninByMobile(ctx, request)
    default:
        response, err = nil, errs.InvalidSigninType
    }

    if err != nil {
        util.HandleGrpcError(c, err)
        return
    }

    if response == nil {
        c.JSON(http.StatusInternalServerError, rest.InternalError())
        return
    }

    userInfo := response.UserInfo
    jwt := middleware.NewJWT()
    token, expiresAt, err := jwt.CreateJWTToken(userInfo.Id, userInfo.Username)

    if err != nil {
        zap.L().Error("create jwt token failed", zap.Error(err))
        c.JSON(http.StatusInternalServerError, rest.InternalError())
        return
    }

    result := &resp.SigninResult{
        UserInfo:         response.UserInfo,
        OrganizationList: response.OrganizationList,
        TokenInfo: &resp.TokenInfo{
            Token:     token,
            ExpiresAt: expiresAt,
        },
    }
    c.JSON(http.StatusOK, rest.Success(result))
}
