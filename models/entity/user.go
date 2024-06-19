package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index , column:deleted_at"` // `json:"-" gorm:"index , column:deleted_at"` (untuk menghilangkan data array di json)
}
