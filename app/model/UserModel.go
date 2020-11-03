package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	password      string    `json:"password"`
	Email       string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
