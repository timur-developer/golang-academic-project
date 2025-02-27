package main

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName string `json:"task_name"`
	IsDone   bool   `json:"is_done"`
}
