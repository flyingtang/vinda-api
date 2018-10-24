package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"vinda-api/conf"
	"vinda-api/models"
	"vinda-api/routers"
)

var r *gin.Engine

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))
	os.Chdir(apppath)
}

func TestMain(m *testing.M) {
	gin.SetMode("test")
	conf.New()
	models.New()
	r = routers.New()
	m.Run()
}

func TestCreateCategory(t *testing.T) {

	//only test good
	Convey("test Category", t, func() {
		const url = "/api/v1/category"

		Convey("should be create success", func() {

			b := map[string]string{
				"name":        "go学习",
				"description": "just a desc about go study",
			}
			d, _ := json.Marshal(b)

			req := httptest.NewRequest("POST", url, bytes.NewReader(d))
			req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
			body := w.Body
			fmt.Println(body)
		})
		Convey("should be find all categories", func() {
			req := httptest.NewRequest("GET", url, nil)
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			So(res.Code, ShouldEqual, http.StatusOK)
			//body := res.Body
			//fmt.Println(body)
		})
	})
}
