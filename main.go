package main

import (
	"vinda-api/conf"
	_ "vinda-api/conf"
	"vinda-api/models"
	"vinda-api/routers"
)

func main() {
	//gin.SetMode("test")
	// Initial config
	conf.New()

	// Initial database
	db := models.New()

	defer db.Close()

	// Initial router
	r := routers.New()
	serverUrl := conf.GlobalConfig.GetHttpAddrPort()
	//
	r.Run(serverUrl)
}
