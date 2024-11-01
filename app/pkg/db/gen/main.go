package main

import (
    "colab/app/pkg/data/model"
    "colab/app/pkg/data/querier"
    "colab/pkg/config"
    "colab/pkg/resource"
    "gorm.io/gen"
)

func main() {
    g := gen.NewGenerator(gen.Config{
        OutPath:           "./app/pkg/data/dal",
        OutFile:           "base.go",
        ModelPkgPath:      "./app/pkg/data/model",
        Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
        FieldWithIndexTag: true,
        FieldWithTypeTag:  true,
        FieldNullable:     true,
        FieldCoverable:    true,
        FieldSignable:     true,
    })

    mysqlCfg := &config.MysqlConfig{
        Username: "root",
        Password: "password",
        Host:     "127.0.0.1",
        Port:     3306,
        Database: "colab",
    }
    db := resource.NewDBConn(mysqlCfg)
    g.UseDB(db)
    g.GenerateAllTable()
    g.ApplyBasic(
        &model.Menu{},
        &model.Organization{},
        &model.OrganizationMember{},
    )
    g.ApplyInterface(func(querier.UserQuerier) {}, model.User{})
    g.ApplyInterface(func(querier.TaskQuerier) {}, model.Task{})
    g.Execute()

}
