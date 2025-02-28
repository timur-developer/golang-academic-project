package taskService

import (
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	if task.TaskName != "" {
		err := r.db.Model(&Task{}).Where("id = ?", id).Update("task_name", task.TaskName).Error
		if err != nil {
			return Task{}, err
		}
		return task, nil
	}
	return Task{}, fmt.Errorf("Could not update task")
}

func (r *taskRepository) DeleteTaskByID(id uint) (Task, error) {
	if id != 0 {
		var deletedTask Task
		r.db.First(&deletedTask, "id = ?", id)
		err := r.db.Delete(&Task{}, id).Error
		if err != nil {
			return Task{}, err
		}
		return deletedTask, nil
	}
	return Task{}, fmt.Errorf("There is no task with such ID")
}
