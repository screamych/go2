package handlers

import (
	"SecondHomework/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"
)

type MakeTask struct {
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type ResponseCreated struct {
	Result bool `json:"result"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/task/" {
		if r.Method == http.MethodPost {
			createTaskHandler(w, r)
		} else if r.Method == http.MethodGet {
			getAllTasksHandler(w, r)
		} else if r.Method == http.MethodDelete {
			deleteAllTasksHandler(w, r)
		} else {
			http.Error(
				w,
				fmt.Sprintf("expect method GET, POST, DELETE at '/task', got %v", r.Method),
				http.StatusMethodNotAllowed,
			)
			return
		}
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Task create handling at %s\n", r.URL.Path)

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if mediaType != "application/json" {
		http.Error(
			w,
			"expect application/json Content-Type",
			http.StatusUnsupportedMediaType,
		)
	}

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	var mt MakeTask
	if err := jsonDecoder.Decode(&mt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tagsList := strings.Join(mt.Tags, ",")
	result, err := models.CreateTask(mt.Text, tagsList, mt.Due)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(ResponseCreated{
		Result: result,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get all tasks handling at %s\n", r.URL.Path)

	allTasks, err := models.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResp, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func deleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete all tasks handling at %s\n", r.URL.Path)

	result, err := models.DeleteAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(ResponseCreated{
		Result: result,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
