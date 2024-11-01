package util

import (
    "colab/app/colab-api/internal/global"
    "colab/app/pkg/model/rest"
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.uber.org/zap"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "net/http"
    "strings"
)

func removeTopStruct(fields map[string]string) map[string]string {
    resp := make(map[string]string)
    for field, err := range fields {
        resp[field[strings.Index(field, ".")+1:]] = err
    }
    return resp
}

func HandleValidatorError(c *gin.Context, err error) {
    var errs validator.ValidationErrors
    var errMsg any

    if errors.As(err, &errs) {
        errMsg = removeTopStruct(errs.Translate(*global.Translator))
    } else {
        errMsg = err.Error()
    }

    c.JSON(http.StatusBadRequest, gin.H{
        "msg": errMsg,
    })
}

func HandleGrpcError(c *gin.Context, err error) {
    s, ok := status.FromError(err)
    if !ok {
        zap.L().Error("unknown error, error", zap.Error(err))
        c.JSON(http.StatusInternalServerError, rest.InternalError())
        return
    }
    code := uint32(s.Code())
    switch s.Code() {
    case codes.AlreadyExists, codes.InvalidArgument:
        c.JSON(http.StatusBadRequest, rest.Fail(code, s.Message()))
    case codes.Unauthenticated:
        c.JSON(http.StatusUnauthorized, rest.Fail(code, s.Message()))
    case codes.PermissionDenied:
        c.JSON(http.StatusForbidden, rest.Fail(code, s.Message()))
    case codes.NotFound:
        c.JSON(http.StatusNotFound, rest.Fail(code, s.Message()))
    case codes.Internal:
        zap.S().Errorf("internal error, msg: %s", s.Message())
        c.JSON(http.StatusInternalServerError, rest.Fail(code, "internal error"))
    default:
        zap.S().Errorf("other error, code: %d, msg: %s", s.Code(), s.Message())
        c.JSON(http.StatusInternalServerError, rest.Fail(code, "internal error"))
    }
    return

}
