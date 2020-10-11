package models

import "gorm.io/gorm"

// User user model
type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Token     string `gorm:"-"`
	Password  string
}
