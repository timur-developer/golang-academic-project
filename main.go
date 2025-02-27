package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Structure to work with request
type requestBody struct {
	ID       uint   `json:"id"`
	TaskName string `json:"task_name"`
	IsDone   bool   `json:"is_done"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{}
	if err := DB.Find(&tasks).Error; err != nil {
		log.Fatal("Could not get tasks")
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&tasks); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	req := requestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Can not decode request", http.StatusBadRequest)
	}
	if req.TaskName == "" {
		http.Error(w, "Task-message could not be empty", http.StatusBadRequest)
	}
	task := Task{
		TaskName: req.TaskName,
		IsDone:   req.IsDone,
	}
	if err := DB.Create(&task).Error; err != nil {
		http.Error(w, "Could not create task", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Task created successfully")
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	req := requestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Can not decode request", http.StatusBadRequest)
	}

	if req.TaskName != "" {
		if err := DB.Model(&Task{}).Where("id = ?", req.ID).Update("task_name", req.TaskName).Error; err != nil {
			http.Error(w, "Could not update task", http.StatusBadRequest)
		}
	}
	//if err := DB.Model(&Task{}).Where("id = ?", req.ID).Update("is_done", req.IsDone).Error; err != nil {
	//	http.Error(w, "Could not update task", http.StatusBadRequest)
	//}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated successfully")
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	req := requestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Could not decode request", http.StatusBadRequest)
	}
	if req.ID != 0 {
		if err := DB.Delete(&Task{}, req.ID).Error; err != nil {
			http.Error(w, "Could not delete task", http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task deleted successfully")
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")
	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/messages", DeleteHandler).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Could not run server")
	}
}
