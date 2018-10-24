package models

const ArticleSchema = `
  create table if not exists Article(
		id int primary key auto_increment,
		title varchar(255) ,
		description text,
		status int default 1,
		content longtext not null,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP 
		categoryId int,
		constraint index(title)
		constraint fk_article_category foreign key (categoryId) references Category(id)
	)`
