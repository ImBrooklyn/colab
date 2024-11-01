package handler

import (
    "colab/api/common"
    "colab/api/grpc/user"
    "colab/app/colab-user/internal/convert"
    "colab/app/pkg/data/dal"
    "colab/app/pkg/errs"
    "context"
    "errors"
    "gorm.io/gorm"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
    return &UserServer{}
}

func (u *UserServer) GetUserInfo(ctx context.Context, request *common.IDRequest) (*user.UserInfo, error) {
    q := dal.User
    userModel, err := q.WithContext(ctx).Where(q.ID.Eq(request.Id)).First()
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errs.UserNotExist
    }
    response := convert.UserModelToUserInfo(userModel)
    return response, nil
}
