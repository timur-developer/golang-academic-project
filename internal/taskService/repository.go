package taskService

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error)
	DeleteTaskByID(id uint) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, fmt.Errorf("Could not create the task: %v", err)
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return []Task{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Could not get tasks: %v", err))
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error) {
	if err := r.db.Model(&Task{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return Task{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("Could not update the task: %v", err))
	}
	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no such task: %v", err))
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) (Task, error) {
	if id != 0 {
		var deletedTask Task
		if err := r.db.First(&deletedTask, id).Error; err != nil {
			return Task{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no task with such ID: %v", err))
		}
		if err := r.db.Delete(&Task{}, id).Error; err != nil {
			return Task{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Could not delete the task: %v", err))
		}
		return deletedTask, nil
	}
	return Task{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no task with such ID"))
}
