package accounts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c * gin.Context){
	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
}

func Signup(c *gin.Context){

}

func Find(c *gin.Context){
	fmt.Println("Find")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
	})
}

func FindOne(c *gin.Context){

}
