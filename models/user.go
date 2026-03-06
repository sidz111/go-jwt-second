package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
