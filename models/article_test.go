package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"vinda-api/conf"
)

func TestMain(m *testing.M) {
	gin.SetMode("test")
	os.Chdir("../")
	conf.New()
	// Initial database
	db := New()

	defer db.Close()
	m.Run()
}

func TestArticle_CreateArticle(t *testing.T) {

	Convey("test create  article", t, func() {
		Convey("should insert success", func() {
			var a = Article{
				//Title: "这是title",
				//Content:"这是内容",
				CategoryId: 1,
			}
			err := CreateArticle(&a)
			fmt.Println(err)
			So(err, ShouldBeNil)
		})
		Convey("should insert fail", func() {
			var a = Article{
				//Title: "这是title",
				//Content:"这是内容",
				//CategoryId: 1,
			}
			err := CreateArticle(&a)
			fmt.Println(err)
			So(err, ShouldBeError)
		})
	})
}

func TestFindArticle(t *testing.T) {
	Convey("test find article", t, func() {
		Convey("should find success", func() {

			as, err := FindArticle(1)

			So(err, ShouldBeNil)
			t.Log(as, "as")
		})
		Convey("should find success default page 1", func() {

			as, err := FindArticle(-11)

			So(err, ShouldBeNil)
			t.Log(as, "as")
		})
	})
}
