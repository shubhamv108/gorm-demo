package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Id   int32
	Name string
	Contact Contact `gorm:"foreignKey:PersonRefer;references:Id"`
}