package domain

import (
    "colab/api/grpc/user"
    "colab/app/pkg/data/dal"
    "colab/app/pkg/errs"
    "context"
    "go.uber.org/zap"
)

type OrganizationDomain struct {
}

func NewOrganizationDomain() *OrganizationDomain {
    return &OrganizationDomain{}
}

func (d *OrganizationDomain) GetOrganizationListByUserID(ctx context.Context, userID int64) ([]*user.OrganizationInfo, error) {
    o := dal.Organization
    om := dal.OrganizationMember
    var organizations []*user.OrganizationInfo
    err := o.WithContext(ctx).Select(o.ALL).
        LeftJoin(om, o.ID.EqCol(om.OrgID)).
        // TODO gorm 未对关联表的软删除加条件
        Where(om.UserID.Eq(userID), om.DeletedAt.IsNull()).
        Scan(&organizations)

    if err != nil {
        zap.L().Error("db error", zap.Error(err))
        return nil, errs.DBError
    }
    
    return organizations, nil
}
