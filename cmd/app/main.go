package main

import (
	"academyProject/internal/database"
	"academyProject/internal/handlers"
	"academyProject/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.Db)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostHandler).Methods("POST")
	router.HandleFunc("/api/patch", handler.PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete", handler.DeleteHandler).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Could not run server")
	}
}
