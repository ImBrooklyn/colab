package initialize

import (
    "colab/app/colab-user/conf"
    "colab/app/colab-user/internal/global"
    "colab/app/pkg/data/dal"
    "colab/pkg/resource"
)

func InitMiddlewares(cfg *conf.UserSrvConfig) {
    global.Redis = resource.NewRedisConn(cfg.Redis)
    db := resource.NewDBConn(cfg.Mysql)
    dal.SetDefault(db)
}
