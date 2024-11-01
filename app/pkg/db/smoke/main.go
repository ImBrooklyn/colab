package main

import (
    "colab/app/pkg/data/dal"
    "colab/app/pkg/data/example"
    "colab/pkg/config"
    "colab/pkg/resource"
    "fmt"
)

func main() {
    mysqlCfg := &config.MysqlConfig{
        Username: "root",
        Password: "password",
        Host:     "127.0.0.1",
        Port:     3306,
        Database: "colab",
    }
    db := resource.NewDBConn(mysqlCfg)
    dal.SetDefault(db)

    q := dal.Task
    result, err := q.QueryTaskByAttr([]example.Attr{
        {Code: "title", Value: "123"},
        {Code: "creator", Value: "huixin.yu"},
    })
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(len(result))
}
