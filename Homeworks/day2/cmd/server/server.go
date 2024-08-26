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

	// TMP tests
	// ---
	// result, err := models.DeleteTaskById(15)
	// if err != nil {
	// return fmt.Errorf("db error %s", err)
	// }
	// log.Println(result)
	// TMP tests

	router := mux.NewRouter()
	router.HandleFunc("/task/", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/task/", handlers.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/task/{id}", handlers.GetTaskHandler).Methods("Get")
	router.HandleFunc("/task/", handlers.DeleteAllTasksHandler).Methods("DELETE")
	return http.ListenAndServe(":8080", router)
}
