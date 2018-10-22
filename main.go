package main

import (
	"github.com/gin-gonic/gin"
	"vinda-api/models"
	"vinda-api/routers"
)


func main(){

	// 初始化数据库
	db := models.InitialDatabse()
	defer db.Close()

	r := gin.Default()
	// 初始化路由
	routers.InitialRouter(r)

	//
	r.Run("0.0.0.0:3000")
}