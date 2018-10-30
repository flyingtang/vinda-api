package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"vinda-api/conf"
)

var schemas = []string{
	`
	 create table if not exists tb_account(
		id int primary key auto_increment,
		username varchar(255) unique  not null,
		password char(128) not null,
		enabled tinyint(1) default 1,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP 
	);`,
	`
	create table if not exists tb_category(
		id int primary key auto_increment,
		name varchar(255) not null not null,
		description text,
		enabled tinyint default 1,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP 
	);`, `
	create table if not exists tb_article
	(
		id int primary key auto_increment,
		title varchar(255) not null,
		description text,
		status int default 1,
		content longtext not null,
		markdown longtext not null,
		main_pic varchar(255),
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
		category_id int,
		index(title),
		constraint fk_article_category foreign key (category_id) references tb_category(id)
	)`}
var globalDB *sqlx.DB

func New() *sqlx.DB {
	//const url = "root:tt1234@tcp(ixiaotang.cn:3307)/vinda?charset=utf8&parseTime=True&loc=Local"
	url, err := conf.GlobalConfig.GetMySQLUrl()
	if err != nil {
		panic(err.Error())
	}
	url += "?charset=utf8&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	tx := db.MustBegin()
	for i := 0; i < len(schemas); i++ {
		tx.MustExec(schemas[i])
	}
	err = tx.Commit()
	if err != nil {
		panic("initial database table error")
	}
	globalDB = db
	initialAccount()
	return db
}

func initialAccount() {
	const p = "1234"
	const username = "admin"
	password, err := HashPassword(p)

	if err != nil {
		panic("initial account err.Error()")
	}

	const sql = "insert into tb_account (username, password) select ?, ? where not exists (select username from tb_account where username=?)"
	_, err = globalDB.Exec(sql, username, password, username)
	if err != nil {
		panic("globalDB.Exec(sql, username, password, username)" + err.Error())
	}
}
