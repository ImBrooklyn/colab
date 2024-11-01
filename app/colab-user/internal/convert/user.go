package convert

import (
    "colab/api/grpc/user"
    "colab/app/pkg/data/model"
)

func UserModelToUserInfo(userModel *model.User) *user.UserInfo {
    return &user.UserInfo{
        Id:       userModel.ID,
        Username: userModel.Username,
        Nickname: userModel.Nickname,
        Mobile:   userModel.Mobile,
        Email:    userModel.Email,
    }
}
