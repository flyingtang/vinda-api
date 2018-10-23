package models

import (
	"cotton/models/account"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSignup(t *testing.T) {
	a := account.Account{
		Username: "xiangang",
		Password: "1234",
	}
	Convey("当用户名不存在的时候")
}

func TestLogin(t *testing.T) {

}
