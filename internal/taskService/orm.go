package taskService

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	TaskName  string    `json:"task_name"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at "json:"updated_at"`
}
