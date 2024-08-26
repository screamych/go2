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
	// task, err := models.GetTaskById(1)
	// if err != nil {
	// return fmt.Errorf("db error %s", err)
	// }
	// fmt.Println(task)
	// fmt.Println()
	// ---
	// result, err := models.DeleteTaskById(15)
	// if err != nil {
	// return fmt.Errorf("db error %s", err)
	// }
	// log.Println(result)
	// TMP tests

	router := mux.NewRouter()
	router.HandleFunc("/task/", handlers.TaskHandler)
	return http.ListenAndServe(":8080", router)
}
