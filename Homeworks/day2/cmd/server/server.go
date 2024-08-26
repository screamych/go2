package server

import (
	"SecondHomework/api/handlers"
	"SecondHomework/internal/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() error {
	err := models.InitDB()
	if err != nil {
		return fmt.Errorf("database inititializing error %s", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/task/", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/task/", handlers.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/task/{id}", handlers.GetTaskHandler).Methods("GET")
	router.HandleFunc("/task/", handlers.DeleteAllTasksHandler).Methods("DELETE")
	router.HandleFunc("/task/{id}", handlers.DeleteTaskHandler).Methods("DELETE")

	return http.ListenAndServe(":8080", router)
}
