package resource

import (
    "colab/pkg/config"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
    "log"
    "os"
    "time"
)

func NewDBConn(cfg *config.MysqlConfig) *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

    var err error
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
            TablePrefix:   "colab_",
        },
        Logger: logger.New(
            log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
            logger.Config{
                SlowThreshold: 2 * time.Second,
                LogLevel:      logger.Info,
                Colorful:      true,
            },
        ),
    })
    if err != nil {
        panic(err)
    }

    return db
}
