package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 35,
		Price:    500,
		Title:    "Margarita",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 20,
		Price:    500.50,
		Title:    "Pepperoni",
	}

	pizza3 := Pizza{
		ID:       3,
		Diameter: 32,
		Price:    500.50,
		Title:    "4Cheeses",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

// Наша модель
type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

// Вспомогательная функция для модели (модельный метод)
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)            // StatusCode для запроса
	json.NewEncoder(writer).Encode(db) // Сериализация + запись в writer
}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(request) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id #:", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := ErrorMessage{Message: "pizza with that id does not exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
