package util

import (
    "colab/app/pkg/constants/keys"
    "context"
    "google.golang.org/grpc/metadata"
    "strconv"
)

func GetUserIDFromContext(ctx context.Context) (uid int64, ok bool) {
    md, got := metadata.FromIncomingContext(ctx)
    if !got {
        return
    }

    strings := md.Get(keys.KCtxUseID)
    if strings == nil || len(strings) <= 0 {
        return
    }

    userIDStr := strings[0]
    if userID, err := strconv.ParseInt(userIDStr, 10, 64); err == nil {
        return userID, true
    }
    return
}

func GetUsernameFromContext(ctx context.Context) (username string, ok bool) {
    md, got := metadata.FromIncomingContext(ctx)
    if !got {
        return
    }

    strings := md.Get(keys.KCtxUsername)
    if strings == nil || len(strings) <= 0 || strings[0] == "" {
        return
    }

    return strings[0], true
}
