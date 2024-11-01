package handler

import (
    "colab/api/grpc/user"
    "colab/app/colab-user/internal/domain"
    "colab/app/pkg/errs"
    "colab/app/pkg/util"
    "context"
)

type OrganizationServer struct {
    orgDomain *domain.OrganizationDomain
}

func NewOrganizationServer() *OrganizationServer {
    return &OrganizationServer{
        orgDomain: domain.NewOrganizationDomain(),
    }
}

func (s *OrganizationServer) GetOrganizationList(ctx context.Context, _ *user.OrganizationListRequest) (*user.OrganizationListResponse, error) {
    userID, ok := util.GetUserIDFromContext(ctx)
    if !ok || userID == 0 {
        return nil, errs.MissingUserInfo
    }

    organizations, err := s.orgDomain.GetOrganizationListByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }

    return &user.OrganizationListResponse{
        List: organizations,
    }, nil
}
