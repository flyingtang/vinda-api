package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
	"vinda-api/conf"
)

type Article struct {
	Id          int       `json:"id"`
	Title       string    `form:"title" binding:"required" json:"title"`
	Description string    `form:"description" json:"description" db:"description"`
	Status      int       `json:"status" db:"status"`
	Content     string    `form:"content" binding:"required" json:"content"`
	Markdown    string    `form:"markdown" binding:"required" json:"markdown"`
	MainPic     string    `form:"mainPic" db:"main_pic"  json:"mainPic"`
	Author      string    `form:"author" binding:"required" json:"author"`	 // 来源作者
	Source      string    `form:"source" binding:"required" json:"source"`   //来源
	CategoryId  int       `form:"categoryId" db:"category_id" json:"categoryId"`
	CreatedAt   time.Time `from:"createdAt" binding:"required" db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `form:"updatedAt" db:"updated_at" json:"updatedAt"`
}

func CreateArticle(a *Article) error {
	const sql = "insert into tb_article (title, description, content, category_id, main_pic, markdown, source, author) values (:title, :description, :content, :category_id, :main_pic, :markdown, :source, :author)"
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

	const sql = "select * from tb_article where status=1 limit  ? offset ?"
	err = globalDB.Select(&as, sql, limit, skip)
	if err != nil {
		return
	}
	const sqltotal = "select count(*) from tb_article where status=1"
	err = globalDB.Get(&total, sqltotal)
	return
}

func FindArticleById(id string) (a Article, err error) {
	const sql = "select * from tb_article where id=?"
	err = globalDB.Get(&a, sql, id)
	return a, err
}

func PatchArticle(id string, a *Article) (err error) {

	const sql = "update tb_article set title=?, description=?, content= ? ,category_id=?, main_pic=?, markdown=?, source=?, author=? where id = ?"
	_, err = globalDB.Exec(sql, a.Title, a.Description, a.Content, a.CategoryId, a.MainPic, a.Markdown, id)
	return err
}

func DeletePatchArticle(ids []int) error {

	if len(ids) == 0 {
		return errors.New("empty ids array in deleting category")
	}
	const sql = "update tb_article set status = 0 where id in (?);"
	query, args, err := sqlx.In(sql, ids)
	query = globalDB.Rebind(query)
	_, err = globalDB.Query(query, args...)
	return err
}

func DeleteArticle(id string) error {

	const sql = "update tb_article set status = 0  where id=?"
	_, err := globalDB.Exec(sql, id)
	return err
}
