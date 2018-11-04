package routers

import (
	"github.com/gin-gonic/gin"
	"vinda-api/controllers"
	"vinda-api/controllers/accounts"
	"vinda-api/controllers/articles"
	"vinda-api/controllers/categories"
)

func New() (r *gin.Engine) {

	r = gin.Default()

	// 设置限制
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Static("/static", "./public")
	version := "/api/v1"

	// not auth
	v := r.Group(version)
	{
		v.POST("/account/login", accounts.Login)
		v.POST("/account/signup", accounts.Signup)

		v.GET("/article", articles.Find)
		v.GET("/article/:id", articles.FindOne)
		v.GET("/category", categories.Find)
	}

	// must auth
	authv := r.Group(version, controllers.Auth)
	{
		authv.POST("/upload", controllers.Upload)

		authv.POST("/article", articles.Create)

		authv.DELETE("/article/:id", articles.Delete)
		authv.PATCH("/article/:id", articles.Patch)
		authv.DELETE("/article", articles.DeleteAll)

		authv.POST("/category", categories.Create)

		authv.DELETE("/category", categories.DeleteAll)
		authv.DELETE("/category/:id", categories.Delete)
		authv.PATCH("/category/:id", categories.Patch)

		// 获取用户信息
		authv.GET("/account/me", accounts.Me)
	}
	return r
}
