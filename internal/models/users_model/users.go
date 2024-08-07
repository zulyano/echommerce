package users_model

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
