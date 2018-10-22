package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


var schema = `
	create Table Account (
		id int primary key,
		username varchar(255) unique  not null,
		password char(64) not null,
		enable tinyint(1) default 1,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP 
	)
`



var globalDB *sqlx.DB

func InitialDatabse() (*sqlx.DB) {
	const url = "root:tt1234@tcp(ixiaotang.cn:3307)/vinda?charset=utf8&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	res := db.MustExec(schema)
	fmt.Println(res)
	globalDB = db
	return db
}
