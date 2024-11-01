package handler

import (
    "colab/api/common"
    "colab/api/grpc/message"
    "colab/app/colab-message/internal/global"
    "colab/app/pkg/constants/sms"
    "context"
    "fmt"
    "go.uber.org/zap"
    "math/rand"
    "strings"
    "time"
)

type MessageServer struct {
}

func NewMessageServer() *MessageServer {
    return &MessageServer{}
}

func generateSmsCode(width int) string {
    numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    var sb strings.Builder
    l := len(numeric)
    for i := 0; i < width; i++ {
        _, _ = fmt.Fprintf(&sb, "%d", numeric[r.Intn(l)])
    }
    return sb.String()
}

func (ms *MessageServer) SendSmsCode(ctx context.Context, request *message.SendSmsCodeRequest) (*common.Flag, error) {
    smsCode := generateSmsCode(6)

    biz := sms.Biz(request.Biz)
    key := biz.Key(request.Mobile)
    zap.S().Infof("saving sms code, key: %s", key)
    _, err := global.Redis.Set(ctx, key, smsCode, 5*time.Minute).Result()
    if err != nil {
        zap.L().Error("save to redis err", zap.Error(err))
        return nil, err
    }

    zap.S().Infof("sending sms code of biz '%s' to [%s]: %s", request.Biz, request.Mobile, smsCode)
    return &common.Flag{Success: true}, nil
}
