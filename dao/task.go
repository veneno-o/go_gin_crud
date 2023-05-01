package dao

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content *string
	Status  int
}
