package dao

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   *string        `json:"content"`
	Status    int            `json:"status"`
}
