package articles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context){

	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

func Find(c *gin.Context){
	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
	})
}

func FindOne(c *gin.Context){

	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
	})
}
