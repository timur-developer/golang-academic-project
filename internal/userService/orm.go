package userService

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Email     string    `gorm:"email" json:"email"`
	Password  string    `gorm:"password" json:"password"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at " json:"updated_at"`
}
