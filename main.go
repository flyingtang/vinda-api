package main

import (
	"github.com/gin-gonic/gin"
	"vinda-api/conf"
	_ "vinda-api/conf"
	"vinda-api/models"
	"vinda-api/routers"
)

func main() {

	// 初始化数据库
	db := models.InitialDatabse()
	defer db.Close()

	r := gin.Default()
	// 初始化路由
	routers.InitialRouter(r)
	serverUrl := conf.GlobalConfig.GetHttpAddrPort()
	//
	r.Run(serverUrl)
}
