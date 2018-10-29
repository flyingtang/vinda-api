package models

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	"vinda-api/conf"
)

type Category struct {
	Id          int       `form:"id" json:"id"`
	Name        string    `form:"name" binding:"required" json:"name"`
	Description string    `form:"description" json:"description"`
	Enabled     bool      `from:"enabled" json:"enabled"`
	CreatedAt   time.Time `from:"createdAt" json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `form:"updatedAt" json:"updatedAt" db:"updated_at"`
}

func CreateCategory(cat *Category) error {
	if cat == nil {
		return errors.New("parameter is nil")
	}
	const sql = `insert into tb_category (name, description) values (:name, :description)`
	_, err := globalDB.NamedExec(sql, *cat)
	return err
}

func FindCategory(page int64) (cats []Category, count int64, err error) {

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
	const sql = "select * from tb_category where enabled=true limit ? offset ?"
	const sqlCount = "select count(*) as total from tb_category where enabled=true"

	err = globalDB.Select(&cats, sql, limit, skip)
	if err != nil {
		return nil, 0, err
	}
	err = globalDB.Get(&count, sqlCount)
	return cats, count, err
}

func PathchCategory(id string, cat *Category) (err error) {

	const sql = "update tb_category set name=?, description=?, enabled=? where id = ?"
	_, err = globalDB.Exec(sql, cat.Name, cat.Description, cat.Enabled, id)
	return err
}

func DeletePatchCategory(ids []int) error {
	if len(ids) == 0 {
		return errors.New("empty ids array in deleting category")
	}
	fmt.Println(ids, "000")
	const sql = "update tb_category set enabled = 0 where id in (?);"
	query, args, err := sqlx.In(sql, ids)
	query = globalDB.Rebind(query)
	_, err = globalDB.Query(query, args...)
	return err
}

func DeleteCategory(id string) error {

	const sql = "update tb_category set enabled =false  where id=?;"
	_, err := globalDB.Exec(sql, id)
	return err
}
