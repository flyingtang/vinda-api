package categories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"vinda-api/models"
)

func Create(c *gin.Context) {

	var cat models.Category
	c.ShouldBind(&cat)
	// 校验
	err := models.CreateCategory(&cat)
	fmt.Println(err, "1212")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "创建失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

func Find(c *gin.Context) {
	var page int64
	var err error

	// TODO 过滤条件查询
	p := c.Query("page")
	if len(p) == 0 {
		page = 1
	} else {
		page, err = strconv.ParseInt(p, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "查询失败",
			})
			c.Abort()
			return
		}
	}
	if cates, count, err := models.FindCategory(page); err != nil {
		logrus.Error("models.FindCategory(page) ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询失败",
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询成功",
			"data":    cates,
			"total":   count,
		})
	}
}
