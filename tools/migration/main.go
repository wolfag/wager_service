package main

import (
	"errors"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"github.com/nguyenhoai890/wager_service/pkg/configuration"
	mysqlUtils "github.com/nguyenhoai890/wager_service/pkg/mysql"
	migrationMysql "github.com/nguyenhoai890/wager_service/tools/migration/mysql"
	"log"
)

func main() {
	uri := mysqlUtils.Uri{
		ServerUri: configuration.MysqlServerUri,
		Params: configuration.MysqlServerUriParams,
		DatabaseName: configuration.MysqlDB,
	}
	err := migrationMysql.InitDB(uri)
	if err != nil {
		panic(err)
	}
	migrate, err := migrationMysql.Init(uri.GetFullUri(), configuration.MigratePath)
	if err != nil {
		panic(err)
	}
	err = migrate.Up()
	if err != nil {
		if errors.Is(err, migrateV4.ErrNoChange) {
			log.Print(err.Error())
		} else {
			panic(err)
		}
	}
}
