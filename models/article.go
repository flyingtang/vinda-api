package models

import (
	"time"
	"vinda-api/conf"
)

type Article struct {
	Id          int       `json:"id"`
	Title       string    `form:"title" binding:"required" json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Content     string    `form:"content" binding:"required" json:"content"`
	CategoryId  int       `form:"categoryId" db:"category_id" json:"category_id"`
	CreatedAt   time.Time `from:"createdAt" binding:"required" db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `form:"updatedAt" db:"updated_at" json:"updatedAt"`
}

func CreateArticle(a *Article) error {
	const sql = "insert into tb_article (title, description, content, category_id) values (:title, :description, :content, :category_id)"
	_, err := globalDB.NamedExec(sql, *a)
	return err
}

func FindArticle(page int64) (as []Article, total int64, err error) {

	var skip int64 = 0
	limit := conf.GlobalConfig.PageLimit
	if limit == 0 {
		limit = 10
	}
	if page > 1 {

		skip = int64(limit) * (page - 1)
	}

	const sql = "select * from tb_article   limit  ? offset ?"
	err = globalDB.Select(&as, sql, limit, skip)
	if err != nil {
		return
	}
	const sqltotal = "select count(*) from tb_article"
	err = globalDB.Get(&total, sqltotal)
	return
}
