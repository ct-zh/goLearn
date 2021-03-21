package model

import "github.com/jinzhu/gorm"

type TestProduct struct {
	gorm.Model
	Code  string
	Price uint
}
