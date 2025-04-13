package models

import "gorm.io/gorm"

// Task — модель задачи
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"` // внешний ключ
}