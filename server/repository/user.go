package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"unique" json:"username"`
	Password  string  `json:"-"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Leaves    []Leave `gorm:"foreignKey:UserID" json:"leaves"`
}
