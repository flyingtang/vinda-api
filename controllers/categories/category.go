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

	p := c.DefaultQuery("page", "1")
	page, err = strconv.ParseInt(p, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询失败",
		})
		c.Abort()
		return
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

func Patch(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	c.ShouldBind(&cat)
	const cNameLen = 4
	if len(cat.Name) < cNameLen {
		logrus.Error("更新分类时，名字长度小于" + string(cNameLen))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "更新失败,只是两个字符",
		})
		c.Abort()
		return
	}
	err := models.PathchCategory(id, &cat)
	if err != nil {
		logrus.Error("models.PathchCategory ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "更新失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
	return
}
type idsForm struct {
	Ids []int `form:"ids"`
}
func DeleteAll(c *gin.Context){
	var ids idsForm
	c.ShouldBind(&ids)
	err := models.DeletePatchCategory(ids.Ids)
	if err != nil {
		logrus.Error("models.DeletePatchCategory(ids)", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "批量删除失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
	return
}
func Delete(c *gin.Context){
	i := c.Param("id")
	id, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		logrus.Error("strconv.ParseInt(i, 10, 32)", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "删除失败",
		})
		c.Abort()
		return
	}
	err = models.DeleteCategory(int(id))
	if err != nil {
		logrus.Error("models.DeleteCategory(id)", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "删除失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
	return
}