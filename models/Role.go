package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	ID     uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name   string `json:"name" gorm:"unique;not null"`
	Status bool   `json:"status" gorm:"default:true"`
}
