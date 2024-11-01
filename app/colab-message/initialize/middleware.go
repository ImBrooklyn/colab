package initialize

import (
    "colab/app/colab-message/conf"
    "colab/app/colab-message/internal/global"
    "colab/app/pkg/data/dal"
    "colab/pkg/resource"
)

func InitMiddlewares(cfg *conf.MessageSrvConfig) {
    global.Redis = resource.NewRedisConn(cfg.Redis)

    db := resource.NewDBConn(cfg.Mysql)
    dal.SetDefault(db)
}
