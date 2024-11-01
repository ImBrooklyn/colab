package middleware

import (
    "colab/app/colab-api/internal/global"
    "colab/app/pkg/constants/keys"
    "colab/app/pkg/errs"
    "colab/app/pkg/model/rest"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "go.uber.org/zap"
    "google.golang.org/grpc/status"
    "net/http"
    "time"
)

type JsonWebToken struct {
    AccessSecret []byte
    Expiry       int
}

type JWTClaims struct {
    UserID   int64
    Username string
    jwt.RegisteredClaims
}

var ignores = make(map[string]any)

func init() {
    list := []string{
        "/colab/user/signin",
        "/colab/user/signup",
        "/colab/message/sms-code",
    }
    for _, item := range list {
        ignores[item] = nil
    }
}

func NewJWT() *JsonWebToken {
    return &JsonWebToken{
        AccessSecret: []byte(global.Config.JWT.AccessSecret),
        Expiry:       global.Config.JWT.Expiry,
    }
}

func (j *JsonWebToken) CreateJWTToken(userID int64, username string) (string, *time.Time, error) {
    now := time.Now()
    maxAge := time.Duration(j.Expiry) * 24 * time.Hour
    expiresAt := now.Add(maxAge)
    claims := JWTClaims{
        UserID:   userID,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            NotBefore: jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(expiresAt),
            Issuer:    "colab",
        },
    }
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := jwtToken.SignedString(j.AccessSecret)
    return token, &expiresAt, err
}

func (j *JsonWebToken) parseToken(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
        return j.AccessSecret, nil
    })

    if token == nil || !token.Valid {
        return nil, errs.TokenInvalid
    }

    if err != nil {
        var ve *jwt.ValidationError
        if errors.As(err, &ve) {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, errs.TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, errs.TokenExpiry
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, errs.TokenNotActive
            }
            return nil, errs.TokenInvalid
        }
    }

    if claims, ok := token.Claims.(*JWTClaims); ok {
        return claims, nil
    } else {
        fmt.Println("ok:", ok)
    }
    return nil, errs.TokenInvalid
}

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestURI := c.Request.RequestURI
        if _, ok := ignores[requestURI]; ok {
            zap.S().Infof("jwt auth - ignore uri: %s", requestURI)
            c.Next()
            return
        }

        token := c.Request.Header.Get(keys.KAuthToken)
        if token == "" {
            s, _ := status.FromError(errs.NotSignedIn)
            c.JSON(http.StatusUnauthorized, rest.Fail(uint32(s.Code()), s.Message()))
            c.Abort()
            return
        }

        j := NewJWT()
        claims, err := j.parseToken(token)
        if err == nil {
            c.Set(keys.KCtxUseID, claims.UserID)
            c.Set(keys.KCtxUsername, claims.Username)
            c.Next()
            return
        }

        if errors.Is(err, errs.TokenExpiry) {
            s, _ := status.FromError(err)
            c.JSON(http.StatusUnauthorized, rest.Fail(uint32(s.Code()), s.Message()))
            c.Abort()
            return
        } else {
            s, _ := status.FromError(errs.NotSignedIn)
            c.JSON(http.StatusUnauthorized, rest.Fail(uint32(s.Code()), s.Message()))
            c.Abort()
            return
        }
    }
}
