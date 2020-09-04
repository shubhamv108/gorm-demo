package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Id int32
	MobileNumber string
	EmailId string
	PersonRefer int32
}
