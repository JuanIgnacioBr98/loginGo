package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name     string `json:"name" gorm:"size:100;unique;not null"`
	Email    string `json:"email" gorm:"size:100;unique;not null"`
	Password string `json:"password" gorm:"default:true;not null"`
	RoleID   uint64 `json:"role_id" gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.Name = html.EscapeString(strings.ToUpper(strings.TrimSpace(user.Name)))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
}
