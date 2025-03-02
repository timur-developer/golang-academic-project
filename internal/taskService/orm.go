package taskService

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	TaskName  string    `gorm:"task_name" json:"task_name"`
	IsDone    bool      `gorm:"is_done" json:"is_done"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at " json:"updated_at"`
	UserID    uint      `gorm:"user_id" json:"user_id"`
}
