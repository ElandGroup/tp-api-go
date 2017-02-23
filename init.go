package main

import (
	"net/http"
	"tp-api-go/api"
	"tp-api-go/config"
	"tp-api-go/dao"

	"github.com/labstack/echo"
)

func init() {
	config.InitConfig("config/config.json")
	dao.InitDB("mssql", config.Config.Sample.Conn) //mssql,mysql
	//model.InitMssql("adodb", config.Config.Sample.Conn)

}

func InitApi(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello api")
	})

	v1 := e.Group("/v1")
	v1.POST("/ExpressOrder", api.CreateExpressOrder)

}
