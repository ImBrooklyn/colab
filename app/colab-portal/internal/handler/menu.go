package handler

import (
    "colab/api/grpc/portal"
    "colab/app/colab-portal/internal/convert"
    "colab/app/colab-portal/internal/global"
    "colab/app/pkg/data/dal"
    "colab/app/pkg/errs"
    "context"
    "encoding/json"
    "go.uber.org/zap"
    "time"
)

type MenuServer struct {
}

func NewMenuServer() *MenuServer {
    return &MenuServer{}
}

func (m *MenuServer) GetUserMenus(ctx context.Context, _ *portal.UserMenuRequest) (*portal.UserMenuResponse, error) {
    q := dal.Menu

    redis := global.Redis
    key := "user-menus"
    cache, err := redis.Get(ctx, key).Result()
    if err != nil {
        menuModels, err := q.WithContext(ctx).Order(q.Pid.Asc(), q.Sort.Asc(), q.ID.Asc()).Find()
        if err != nil {
            zap.L().Error("db error", zap.Error(err))
            return nil, errs.DBError
        }
        menus := convert.MenuModelsToMenuItems(menuModels)
        if menuBytes, err := json.Marshal(menus); err != nil {
            zap.L().Error("marshal data error", zap.Error(err))
        } else if _, err := redis.Set(ctx, key, string(menuBytes), 7*24*time.Hour).Result(); err != nil {
            zap.L().Error("save to redis err", zap.Error(err))
        }

        return &portal.UserMenuResponse{Menus: menus}, nil
    }

    var menus []*portal.UserMenuItem
    if err = json.Unmarshal([]byte(cache), &menus); err != nil {
        return nil, errs.DataParsingError
    }

    return &portal.UserMenuResponse{Menus: menus}, nil
}
