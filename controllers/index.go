package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"
	"vinda-api/conf"
)

func Auth(c *gin.Context) {
	var tokenString string
	tokenString = c.GetHeader("authorization")
	if len(tokenString) == 0 || tokenString == "undefined"{
		tokenString = c.Query("authorization")
		if len(tokenString) == 0 || tokenString == "undefined"{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请登录",
			})
			c.Abort()
			return
		}
	}
	jwtSecret := conf.GlobalConfig.JwtSecret
	if len(jwtSecret) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器参数错误,请稍后再试",
		})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		logrus.Errorf("jwt.Parse err %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请重启登录",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		c.Set("user", map[string]interface{}{
			"userId": claims["userId"],
			"username": claims["username"],
		})
		fmt.Println(c.Get("user"))
	} else {
		logrus.Errorf("token.Claims err %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请重启登录",
		})
		c.Abort()
		return
	}
}


func Upload(c *gin.Context){
	file, _ := c.FormFile("file")
	r, err := GetRandomNumber()
	if err != nil {
		logrus.Errorf("random string get error %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败,请稍后再试",
		})
		c.Abort()
		return
	}
	ext := path.Ext(file.Filename)
	fileName := r+ext //重命名
	basePublic := conf.GlobalConfig.BasePublic // TODO 应该校验
	fp := filepath.Join(basePublic, fileName)
	err = c.SaveUploadedFile(file, fp)
	if err != nil {
		logrus.Errorf("c.SaveUploadedFile err %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败,请稍后再试",
		})
		c.Abort()
		return
	}

	// 拼接图片地址
	host := conf.GlobalConfig.PublicHost // TODO 应该校验
	url := filepath.Join(host, "static", fileName)
	fmt.Println(url)
	c.JSON(http.StatusOK, gin.H{
		"message": "上传成功",
		"url": url,
	})
}

// 用md5生成一个随机字符串
func GetRandomNumber()(string, error){
	h:= md5.New()
	t := time.Now().Unix()
	i:= rand.Int63()
	is := strconv.FormatInt(i, 10)
	_, err := h.Write([]byte(strconv.FormatInt(t, 10)))
	return hex.EncodeToString(h.Sum([]byte(is))), err
}