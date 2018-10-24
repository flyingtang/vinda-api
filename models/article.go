package models

import "time"

type Article struct {
	Id          int
	Title       string
	Description string
	Status      int
	Content     string
	CategoryId  int
	CreatedAt   time.Time `from:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `form:"updatedAt" db:"updated_at"`
}
