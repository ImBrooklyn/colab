package initialize

import (
    "colab/app/colab-portal/conf"
    "colab/app/colab-portal/internal/global"
    "colab/app/pkg/data/dal"
    "colab/pkg/resource"
)

func InitMiddlewares(cfg *conf.PortalSrvConfig) {
    global.Redis = resource.NewRedisConn(cfg.Redis)
    db := resource.NewDBConn(cfg.Mysql)
    dal.SetDefault(db)
}
