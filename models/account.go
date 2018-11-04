package models

import (
	"crypto/sha256"
	"encoding/hex"
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
	const sql = "select * from tb_account where username=? limit 1"
	err := globalDB.Get(&a, sql, username)
	return &a, err
}

func Signup(c *Account) error {

	if len(c.Password) < 6 {
		return errors.New("密码不能少于6位")
	}
	// 对密码转码
	h := sha256.New()
	password := h.Sum([]byte(c.Password))
	c.Password = hex.EncodeToString(password)

	const sql = "insert into tb_account (username, password) values (:Username, :Password)"
	_, err := globalDB.NamedExec(sql, *c)
	return err
}

func HashPassword(pass string) (string, error) {
	if len(pass) == 0 {
		return "", errors.New("empty password")
	}
	h := sha256.New()
	password := h.Sum([]byte(pass))
	return hex.EncodeToString(password), nil
}


