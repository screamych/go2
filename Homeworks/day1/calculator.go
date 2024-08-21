package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func OperandSetter(vn string) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	err := os.Setenv(vn+"_VALUE_ENV", strconv.Itoa(rand.Intn(100)))
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(vn + "_VALUE_ENV")
}

func CalcOperands(method string) string {
	first, _ := strconv.Atoi(os.Getenv("FIRST_VALUE_ENV"))
	second, _ := strconv.Atoi(os.Getenv("SECOND_VALUE_ENV"))

	switch method {
	case "/add":
		return fmt.Sprintf("%d + %d = %d", first, second, first+second)
	case "/sub":
		return fmt.Sprintf("%d - %d = %d", first, second, first-second)
	case "/mul":
		return fmt.Sprintf("%d * %d = %d", first, second, first*second)
	case "/div":
		return fmt.Sprintf("%d / %d = %d", first, second, first/second)
	}

	return ""
}

func OperandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/first":
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, OperandSetter("FIRST"))
	case "/second":
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, OperandSetter("SECOND"))
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, CalcOperands(r.URL.Path))
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(
		w,
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"/first  define first operand",
		"/second define second operand",
		"/add    adds two operands",
		"/sub    subtracts second operand from the first",
		"/mul    multiplies both operands",
		"/div    divides the numerator by the denominator",
		"/info   api information",
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/first", OperandsHandler)
	mux.HandleFunc("/second", OperandsHandler)
	mux.HandleFunc("/add", CalcHandler)
	mux.HandleFunc("/sub", CalcHandler)
	mux.HandleFunc("/mul", CalcHandler)
	mux.HandleFunc("/div", CalcHandler)
	mux.HandleFunc("/info", InfoHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
