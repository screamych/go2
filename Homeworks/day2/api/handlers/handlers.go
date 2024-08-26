package handlers

import (
	"SecondHomework/internal/models"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type MakeTask struct {
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type ResponseCreated struct {
	Result bool `json:"result"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Task create handling at %s\n", r.URL.Path)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get all tasks handling at %s\n", r.URL.Path)
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get task by id handling at %s\n", r.URL.Path)
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
	}

	task, err := models.GetTaskById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete all tasks handling at %s\n", r.URL.Path)
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete task by id handling at %s\n", r.URL.Path)
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
	}

	_, err = models.DeleteTaskById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(ResponseCreated{Result: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
