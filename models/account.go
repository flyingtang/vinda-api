package models

import "time"



type Account struct {
	Id uint
	Username string
	password string
	Enable  bool
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
} 