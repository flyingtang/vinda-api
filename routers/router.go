package routers

import (
	"github.com/gin-gonic/gin"
	"vinda-api/controllers/accounts"
	"vinda-api/controllers/articles"
	"vinda-api/controllers"
)

func InitialRouter(r *gin.Engine){

	version := "/api/v1"
	v := r.Group(version)
	{
		// 用户相关路由
		account := v.Group("/account")
		account.POST("/login", accounts.Login)
		account.POST("/signup", accounts.Signup)

		// 下面更认证过得路由
		account.Use(controllers.Auth)
		account.GET("/", accounts.Find)

		article := v.Group("/article", controllers.Auth)
		{
			article.POST("/", articles.Create)
			article.GET("/", articles.Find)
			article.GET("/:id", articles.FindOne)
		}
	}
}