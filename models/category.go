package models

import (
	"errors"
	"github.com/sirupsen/logrus"
	"time"
	"vinda-api/conf"
)

const CategorySchema = `
create table if not exists Category(
		id int primary key,
		name varchar(255) not null not null,
		description text,
		enabled tinyint default 1,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP 
	);
`

type Category struct {
	Id          string    `form:"id"`
	Name        string    `form:"name" binding:"required" json:"name"`
	Description string    `form:"description" json:"description"`
	Enabled     bool      `from:"enabled"`
	CreatedAt   time.Time `from:"createdAt" `
	UpdatedAt   time.Time `form:"updatedAt"`
}

func CreateCategory(cat *Category) error {
	if cat == nil {
		return errors.New("parameter is nil")
	}
	const sql = `insert into tb_category (name, description) values (:name, :description)`
	_, err := globalDB.NamedExec(sql, *cat)
	return err
}

func FindCategory(page int64) (cats *[]Category, count int64, err error) {

	if page > 1 {
		page -= 1
	} else {
		page = 0
	}
	limit := conf.GlobalConfig.PageLimit
	if limit < 10 {
		limit = 10
		logrus.Warnf("PageLimit not find in config, default 1")
	}
	skip := page * int64(limit)
	const sql = "select * from tb_category where limit ? skip ?"
	const sqlCount = "select count(*) from tb_category"

	err = globalDB.Select(cats, sql, limit, skip)
	if err != nil {
		return nil, 0, err
	}
	err = globalDB.Select(&count, sqlCount)
	return cats, count, err
}
