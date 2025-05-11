package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Name   string `json:"name" gorm:"unique;not null"`
	Status bool   `json:"status" gorm:"default:true"`
	Users  []User `json:"users" gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string {
	return "roles"
}
