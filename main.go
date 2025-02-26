package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	req := requestBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprintln(w, "Error decoding request")
	}
	if req.Message != "" {
		task = req.Message
		fmt.Fprintln(w, "Task was updated successfully")
	} else {
		fmt.Fprintln(w, "Task-message could not be empty")
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetHandler).Methods("GET")
	router.HandleFunc("/", PostHandler).Methods("POST")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Could not run server")
	}
}
