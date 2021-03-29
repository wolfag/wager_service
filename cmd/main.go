package main

import (
	"github.com/nguyenhoai890/wager_service/internal/repository"
	service "github.com/nguyenhoai890/wager_service/internal/service"
	"github.com/nguyenhoai890/wager_service/pkg/configuration"
	mysqlUtils "github.com/nguyenhoai890/wager_service/pkg/mysql"
	"github.com/nguyenhoai890/wager_service/pkg/restful"
	"net/http"
)
func main(){
	 uri := mysqlUtils.Uri{
		ServerUri:    configuration.MysqlServerUri,
		Params:       configuration.MysqlServerUriParams,
		DatabaseName: configuration.MysqlDB,
	}
	repoWage, err := repository.Init(uri)
	if err != nil {
		panic(err)
	}
	serviceWager := service.Init(repoWage)
	r := restful.InitGin(serviceWager)
	http.ListenAndServe(":8080", r)
}