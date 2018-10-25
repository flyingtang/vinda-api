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

	version := "/api/v1"

	// not auth
	v := r.Group(version)
	{
		v.POST("/account/login", accounts.Login)
		v.POST("/account/signup", accounts.Signup)
	}

	// must auth
	authv := r.Group(version, controllers.Auth)
	{

		authv.POST("/article", articles.Create)
		authv.GET("/article", articles.Find)
		authv.GET("/article:id", articles.FindOne)

		authv.POST("/category", categories.Create)
		authv.GET("/category", categories.Find)

	}
	return r
}
