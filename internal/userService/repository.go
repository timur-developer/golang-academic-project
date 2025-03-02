package userService

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, updates map[string]interface{}) (User, error)
	DeleteUserByID(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, fmt.Errorf("Could not create the User: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return []User{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Could not get Users: %v", err))
	}
	return users, nil
}

func (r *userRepository) UpdateUserByID(id uint, updates map[string]interface{}) (User, error) {
	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return User{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("Could not update the User: %v", err))
	}
	var updatedUser User
	if err := r.db.First(&updatedUser, id).Error; err != nil {
		return User{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no such User: %v", err))
	}
	return updatedUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) (User, error) {
	if id != 0 {
		var deletedUser User
		if err := r.db.First(&deletedUser, id).Error; err != nil {
			return User{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no User with such ID: %v", err))
		}
		if err := r.db.Delete(&User{}, id).Error; err != nil {
			return User{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Could not delete the User: %v", err))
		}
		return deletedUser, nil
	}
	return User{}, echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("There is no User with such ID"))
}
