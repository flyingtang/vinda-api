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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title或者 content 长度小于6",
		})
		c.Abort()
		return
	}
	// 其他校验
	err := models.CreateArticle(&a)
	if err != nil {
		logrus.Fatalf("a.CreateArticle() err %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
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
	p := c.Query("page")
	if len(p) > 0 {
		if p, err := strconv.ParseInt(p, 10, 64); err == nil {
			page = p
		}
	}
	as, total, err := models.FindArticle(page)
	if err != nil {
		logrus.Error("models.FindArticle(page) err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查找失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
		"data":    as,
		"total":   total,
	})
	return
}

func FindOne(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		logrus.Error("article find one no valid id")
		c.JSON(http.StatusOK, gin.H{
			"message": "查询文章ID无效",
		})
		c.Abort()
		return
	}
	// TODO
	a, err := models.FindArticleById(id)
	if err != nil {
		logrus.Error("models.FindArticleById(id) err", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "查找失败",
		})
		c.Abort()
		return
	}
	fmt.Println("Login")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
		"data":    a,
	})
}

func Patch(c *gin.Context) {
	id := c.Param("id")
	var a models.Article
	c.ShouldBind(&a)
	titleLen := conf.GlobalConfig.TitleLen
	if titleLen < 1 {
		titleLen = 6
	}
	if len(a.Title) < int(titleLen) || len(a.Content) < int(titleLen) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title或者 content 长度小于6",
		})
		c.Abort()
		return
	}
	err := models.PatchArticle(id, &a)
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

func DeleteAll(c *gin.Context) {
	var ids idsForm
	c.ShouldBind(&ids)
	err := models.DeletePatchArticle(ids.Ids)
	if err != nil {
		logrus.Error("models.DeletePatchArticle(ids)", err.Error())
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
func Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		logrus.Error("invalid delete article id")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID 无效",
		})
		c.Abort()
		return
	}
	err := models.DeleteArticle(id)
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
