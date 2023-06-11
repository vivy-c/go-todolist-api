package model

import "gorm.io/gorm"

type Todo struct {
	// ID        int64 `gorm:"primary key;autoIncrement"`
	Todo   string
	Title  string
	Detail string
	gorm.Model
}
