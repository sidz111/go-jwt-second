package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
