package util

import (
    "colab/app/pkg/constants/keys"
    "context"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc/metadata"
    "strconv"
    "time"
)

func NewApiContext(c *gin.Context) (context.Context, context.CancelFunc) {
    ctx := context.Background()
    userID := c.GetInt64(keys.KCtxUseID)
    username := c.GetString(keys.KCtxUsername)
    if userID > 0 && username != "" {
        pairs := metadata.Pairs(
            keys.KCtxUseID, strconv.FormatInt(userID, 10),
            keys.KCtxUsername, username,
        )
        ctx = metadata.NewOutgoingContext(ctx, pairs)
    }

    return context.WithTimeout(ctx, 2*time.Second)
}
