package models

import "time"

type User struct {
	Id        int64
	CreatedAt time.Time `orm:"null"`
	UpdatedAt time.Time `orm:"null"`
	DeletedAt time.Time `orm:"null"`
	Name      string
	Email     string `orm:"unique"`
	Password  string `orm:"unique"`
}
