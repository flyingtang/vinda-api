package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"vinda-api/conf"
)

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("authorization")
	fmt.Println(tokenString, "123")
	if len(tokenString) == 0 || tokenString == "undefined"{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		c.Abort()
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器参数错误,请稍后再试",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		c.Set("user", &map[string]interface{}{
			"userId": claims["userId"],
			"username": claims["username"],
		})
		fmt.Println(c.Get("user"))
	} else {
		logrus.Errorf("token.Claims err %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器参数错误,请稍后再试",
		})
		c.Abort()
		return
	}
}
