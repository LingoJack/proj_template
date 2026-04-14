package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;size:128;not null" json:"email"`
	Password string `gorm:"size:256;not null" json:"-"`
	Role     string `gorm:"size:32;default:'user'" json:"role"`
}
