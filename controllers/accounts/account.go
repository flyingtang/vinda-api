package accounts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"vinda-api/conf"
	"vinda-api/models"
)

// 登录
func Login(c *gin.Context) {
	var account models.Account
	c.ShouldBind(&account)
	account.Username = strings.TrimSpace(account.Username)
	account.Password = strings.TrimSpace(account.Password)

	_account, err := models.Login(account.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	h := sha256.New()
	password := hex.EncodeToString(h.Sum([]byte(account.Password)))

	if password != _account.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
		c.Abort()
		return
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   _account.Id,
		"username": _account.Username,
	})
	secret := conf.GlobalConfig.JwtSecret
	if len(secret) == 0 {
		gin.LoggerWithWriter(os.Stderr, "服务器获取JwtSecret失败，请检查配置文件")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器异常",
		})
		c.Abort()
		return
	}
	s, err := t.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "登录失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   s,
	})
}

// 注册
func Signup(c *gin.Context) {
	var account models.Account
	c.ShouldBind(&account)
	account.Username = strings.TrimSpace(account.Username)
	account.Password = strings.TrimSpace(account.Password)

	if len(account.Username) < 6 || len(account.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名或者密码少于6位",
		})
		c.Abort()
		return
	}
	if err := models.Signup(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

// 查找用户
func Find(c *gin.Context) {
	fmt.Println("Find")
	c.JSON(http.StatusOK, gin.H{
		"message": "查找成功",
	})
}

// 查找一个用户
func FindOne(c *gin.Context) {

}
