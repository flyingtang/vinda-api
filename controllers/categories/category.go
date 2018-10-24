package categories

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	// TODO 过滤条件查询
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询失败",
		})
		c.Abort()
		return
	}

	if cates, count, err := models.FindCategory(page); err != nil {
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
