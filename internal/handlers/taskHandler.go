package handlers

import (
	"academyProject/internal/taskService"
	"academyProject/internal/web/tasks"
	"context"
)

type requestBody struct {
	ID       uint    `json:"id"`
	TaskName *string `json:"task_name"`
	IsDone   *bool   `json:"is_done"`
}

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTasks(_ context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	taskRequest := request.Body

	deletedTask, err := h.Service.DeleteTaskByID(*taskRequest.Id)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasks200JSONResponse{
		Id:        &deletedTask.ID,
		TaskName:  &deletedTask.TaskName,
		IsDone:    &deletedTask.IsDone,
		CreatedAt: &deletedTask.CreatedAt,
		UpdatedAt: &deletedTask.UpdatedAt,
	}
	return response, nil
}

func (h *Handler) PatchTasks(_ context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	taskRequest := request.Body

	updates := make(map[string]interface{})
	if taskRequest.TaskName != nil {
		updates["task_name"] = *taskRequest.TaskName
	}
	if taskRequest.IsDone != nil {
		updates["is_done"] = *taskRequest.IsDone
	}

	updatedTask, err := h.Service.UpdateTaskByID(*taskRequest.Id, updates)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasks200JSONResponse{
		Id:        &updatedTask.ID,
		TaskName:  &updatedTask.TaskName,
		IsDone:    &updatedTask.IsDone,
		CreatedAt: &updatedTask.CreatedAt,
		UpdatedAt: &updatedTask.UpdatedAt,
	}

	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:        &tsk.ID,
			TaskName:  &tsk.TaskName,
			IsDone:    &tsk.IsDone,
			CreatedAt: &tsk.CreatedAt,
			UpdatedAt: &tsk.UpdatedAt,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		TaskName: *taskRequest.TaskName,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:        &createdTask.ID,
		TaskName:  &createdTask.TaskName,
		IsDone:    &createdTask.IsDone,
		CreatedAt: &createdTask.CreatedAt,
		UpdatedAt: &createdTask.UpdatedAt,
	}

	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
