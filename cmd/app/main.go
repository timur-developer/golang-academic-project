package main

import (
	"academyProject/internal/database"
	"academyProject/internal/handlers"
	"academyProject/internal/taskService"
	"academyProject/internal/userService"
	"academyProject/internal/web/tasks"
	"academyProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
