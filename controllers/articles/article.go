package articles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"vinda-api/conf"
	"vinda-api/models"
)

func Create(c *gin.Context) {

	var a models.Article
	c.ShouldBind(&a)
	titleLen := conf.GlobalConfig.TitleLen
	if titleLen < 1 {
		titleLen = 6
	}

	if len(a.Title) < int(titleLen) || len(a.Content) < int(titleLen) {
		c.JSON(http.StatusOK, gin.H{
			"message": "title或者 content 长度小于6",
		})
		c.Abort()
		return
	}
	// 其他校验
	err := models.CreateArticle(&a)
	if err != nil {
		logrus.Fatalf("a.CreateArticle() err %s", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "创建失败,请检查参数传递",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

func Find(c *gin.Context) {
	var page int64 = 1
	p:= c.Query("page")
	if len(p) > 0 {
		if p , err := strconv.ParseInt(p, 10, 64); err == nil {
			page = p
		}
	}
	as,total, err := models.FindArticle(page)
	if err != nil {
		logrus.Error("models.FindArticle(page) err", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "查找失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
		"data": as,
		"total": total,
	})
	return
}

func FindOne(c *gin.Context) {

	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
	})
}
