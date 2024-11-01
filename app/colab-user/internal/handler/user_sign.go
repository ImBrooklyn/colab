package handler

import (
    "colab/api/common"
    "colab/api/grpc/user"
    "colab/app/colab-user/internal/convert"
    "colab/app/colab-user/internal/domain"
    "colab/app/colab-user/internal/global"
    "colab/app/pkg/constants/sms"
    "colab/app/pkg/data/dal"
    "colab/app/pkg/data/model"
    "colab/app/pkg/errs"
    "colab/pkg/util/password"
    "context"
    "errors"
    "fmt"
    "github.com/go-redis/redis/v8"
    "go.uber.org/zap"
    "gorm.io/gen/field"
    "gorm.io/gorm"
    "strings"
)

type UserSignServer struct {
    orgDomain *domain.OrganizationDomain
}

func NewUserSignServer() *UserSignServer {
    return &UserSignServer{
        orgDomain: domain.NewOrganizationDomain(),
    }
}

func (u *UserSignServer) Signup(ctx context.Context, request *user.SignupRequest) (*common.Flag, error) {
    // sms code
    biz := sms.BizSignup
    key := biz.Key(request.Mobile)
    zap.S().Infof("getting sms codes, key: %s", key)

    result, err := global.Redis.Get(ctx, key).Result()
    if errors.Is(err, redis.Nil) {
        return nil, errs.VerifyCodeNotExist
    } else if err != nil {
        zap.L().Error("redis error when getting sms code", zap.Error(err))
        return nil, errs.RedisError
    } else if result != request.SmsCode {
        return nil, errs.VerifyCodeMismatch
    }

    // username, email, mobile
    q := dal.User
    if count, _ := q.WithContext(ctx).Where(q.Username.Eq(request.Username)).Count(); count != 0 {
        return nil, errs.UsernameExists
    }

    if count, _ := q.WithContext(ctx).Where(q.Email.Eq(request.Email)).Count(); count != 0 {
        return nil, errs.EmailExists
    }

    if count, _ := q.WithContext(ctx).Where(q.Mobile.Eq(request.Mobile)).Count(); count != 0 {
        return nil, errs.MobileExists
    }

    // encode password
    salt, pwd := password.Encode(request.Password, password.DefaultOptions())

    userID, err := global.Snowflake.NextID(ctx, &common.Empty{})
    if err != nil {
        zap.L().Error("error when getting snowflake id", zap.Error(err))
        return nil, errs.RPCError
    }

    orgID, err := global.Snowflake.NextID(ctx, &common.Empty{})
    if err != nil {
        zap.L().Error("error when getting snowflake id", zap.Error(err))
        return nil, errs.RPCError
    }

    userModel := model.User{
        ID:       userID.Id,
        Username: request.Username,
        Nickname: request.Nickname,
        Password: fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, pwd),
        Email:    request.Email,
        Mobile:   request.Mobile,
    }

    organizationModel := model.Organization{
        ID:          orgID.Id,
        Name:        fmt.Sprintf("%s的个人组织", request.Nickname),
        Owner:       userID.Id,
        Avatar:      nil,
        Description: nil,
        Personal:    true,
    }

    organizationMemberModel := model.OrganizationMember{
        OrgID:   orgID.Id,
        UserID:  userID.Id,
        IsOwner: true,
    }

    err = dal.Q.Transaction(func(tx *dal.Query) error {
        if err := tx.User.WithContext(ctx).Create(&userModel); err != nil {
            return err
        }

        if err := tx.Organization.WithContext(ctx).Create(&organizationModel); err != nil {
            return err
        }

        if err := tx.OrganizationMember.WithContext(ctx).Create(&organizationMemberModel); err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        zap.L().Error("db insert failed", zap.Error(err))
        return nil, errs.SignupFailed
    }
    return &common.Flag{Success: true}, nil
}

func (u *UserSignServer) SigninByMobile(ctx context.Context, request *user.SigninByMobileRequest) (*user.SigninResponse, error) {
    biz := sms.BizSignin
    key := biz.Key(request.Mobile)
    result, err := global.Redis.Get(ctx, key).Result()
    if errors.Is(err, redis.Nil) {
        return nil, errs.VerifyCodeNotExist
    } else if err != nil {
        zap.L().Error("redis error when getting sms code", zap.Error(err))
        return nil, errs.RedisError
    } else if result != request.SmsCode {
        return nil, errs.VerifyCodeMismatch
    }

    // deletes redis key
    defer global.Redis.Del(ctx, key)

    q := dal.User
    userModel, err := q.WithContext(ctx).Where(q.Mobile.Eq(request.Mobile)).First()

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errs.UserNotExist
    }

    userInfo := convert.UserModelToUserInfo(userModel)
    return &user.SigninResponse{
        UserInfo: userInfo,
    }, nil
}

func (u *UserSignServer) SigninByPassword(ctx context.Context, request *user.SigninByPasswordRequest) (*user.SigninResponse, error) {
    q := dal.User

    var expr field.Expr
    switch request.SigninType {
    case 1:
        expr = q.Username.Eq(request.Account)
    case 2:
        expr = q.Email.Eq(request.Account)
    default:
        return nil, errs.InvalidSigninType
    }

    userModel, err := q.WithContext(ctx).Where(expr).First()
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errs.UserNotExist
    }

    parts := strings.Split(userModel.Password, "$")
    if ok := password.Match(request.Password, parts[2], parts[3], password.DefaultOptions()); !ok {
        return nil, errs.IncorrectPassword
    }

    userInfo := convert.UserModelToUserInfo(userModel)

    organizations, err := u.orgDomain.GetOrganizationListByUserID(ctx, userModel.ID)
    if err != nil {
        return nil, err
    }

    return &user.SigninResponse{
        UserInfo:         userInfo,
        OrganizationList: organizations,
    }, nil
}
