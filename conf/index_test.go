package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPath(t *testing.T) {

	Convey("different mode should get different file path", t, func() {
		Convey("should get dev file path", func() {
			gin.SetMode(gin.DebugMode)
			cfp := getPath()
			fmt.Println(cfp)
			So(cfp, ShouldEqual, "json/dev.config.json")
		})
		Convey("should get test file path", func() {
			gin.SetMode(gin.TestMode)
			cfp := getPath()
			fmt.Println(cfp)
			So(cfp, ShouldEqual, "json/test.config.json")
		})
		Convey("should get production  file path", func() {
			gin.SetMode(gin.ReleaseMode)
			cfp := getPath()
			fmt.Println(cfp)
			So(cfp, ShouldEqual, "json/pro.config.json")
		})
		//
		//Convey("should get empty file path", func() {
		//	gin.SetMode("xxx")
		//	//cfp := getPath()
		//	//fmt.Println(cfp)
		//	//So(cfp, ShouldEqual, "")
		//
		//})
	})
}

func TestGetConfig(t *testing.T) {
	c := getConfig()
	if c != nil {
		t.Log("成功")
	} else {
		t.Error("失败了")
	}
}

func TestConfig_GetHttpAddrPort(t *testing.T) {
	Convey("test all ", t, func() {

		Convey("should get default server listen url", func() {
			c := new(Config)
			url := c.GetHttpAddrPort()
			So(url, ShouldEqual, "127.0.0.1:3000")
		})
		Convey("should get custom server listen url", func() {
			c := new(Config)
			c.HttpPort = "4000"
			c.HttpAddr = "baidu.com"
			url := c.GetHttpAddrPort()
			So(url, ShouldEqual, "baidu.com:4000")
		})
	})
}

func TestConfig_GetMySQLUrl(t *testing.T) {
	Convey("test all mysql connect url", t, func() {

		Convey("should get error ,because miss username or password or database name", func() {
			Convey("should get error because of missing  username", func() {
				c := new(Config)
				c.MySQL = MySQL{
					Password: "1233",
					Database: "vind",
				}
				url, err := c.GetMySQLUrl()
				So(url, ShouldEqual, "")
				So(err, ShouldBeError)
			})
			Convey("should get error because of missing  password", func() {
				c := new(Config)
				c.MySQL = MySQL{
					Username: "123",
					//Password:"1233",
					Database: "vind",
				}
				url, err := c.GetMySQLUrl()
				So(url, ShouldEqual, "")
				So(err, ShouldBeError)
			})
			Convey("should get error because of missing  database name", func() {
				c := new(Config)
				c.MySQL = MySQL{
					Username: "123",
					Password: "1233",
					//Database: "vind",
				}
				url, err := c.GetMySQLUrl()
				So(url, ShouldEqual, "")
				So(err, ShouldBeError)
			})
		})
		Convey("should be successful using default value eg: Host , port", func() {
			c := new(Config)
			c.MySQL = MySQL{
				Username: "123",
				Password: "1233",
				Database: "vind",
			}
			url, err := c.GetMySQLUrl()
			fmt.Printf(url)
			So(url, ShouldEqual, "123:1233@tcp(127.0.0.1:3306)/vind")
			So(err, ShouldBeNil)
		})
	})
}
