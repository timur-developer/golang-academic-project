package handlers

import (
	"academyProject/internal/taskService"
	"encoding/json"
	"net/http"
)

type requestBody struct {
	ID       uint   `json:"id"`
	TaskName string `json:"task_name"`
}

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&tasks); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Can not decode request", http.StatusBadRequest)
		return
	}
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	var updatedTask taskService.Task
	req := requestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Could not decode request", http.StatusBadRequest)
		return
	}
	updatedTask = taskService.Task{
		ID:       req.ID,
		TaskName: req.TaskName,
	}

	task, err := h.Service.UpdateTaskByID(updatedTask.ID, updatedTask)
	if err != nil {
		http.Error(w, "Could not update task", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Could not decode request", http.StatusBadRequest)
		return
	}
	err := h.Service.DeleteTaskByID(task.ID)
	if err != nil {
		http.Error(w, "Could not delete task", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
