package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type Account struct {
	Id        uint      `form:"id" json:"id"`
	Username  string    `form:"username" json:"username"`
	Password  string    `form:"password" json:"password"`
	Enabled   bool      `form:"enabled" json:"enabled"`
	CreatedAt time.Time `from:"createdAt" db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `form:"updatedAt" db:"updated_at" json:"updatedAt"`
}

func GetAccount(id string) {

}

func Login(username string) (*Account, error) {
	var a Account
	const sql = "select * from Account where username=?"
	err := globalDB.Select(&a, sql, username)
	return &a, err
}

func Signup(c *Account) error {

	if len(c.Password) < 6 {
		return errors.New("密码不能少于6位")
	}
	// 对密码转码
	h := sha256.New()
	password := h.Sum([]byte(c.Password))
	fmt.Println(password, "p")
	c.Password = hex.EncodeToString(password)
	fmt.Println(c.Password, "pp")

	const sql = "insert into tb_ccount (username, password) values (:Username, :Password)"
	_, err := globalDB.NamedExec(sql, *c)
	return err
}
