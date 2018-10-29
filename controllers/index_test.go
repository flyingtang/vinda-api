package controllers

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)
func TestGetRandomNumber(t *testing.T) {
	Convey("test TestGetRandomNumber", t, func() {
		u := []string{
			"1","1","2","3",
		}
		ur := make([]string, 0)
		for range u {
			r, _ := GetRandomNumber()
			ur = append(ur, r)
		}
		fmt.Println(ur)
	})
}
