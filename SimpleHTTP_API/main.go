package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, Linux</h1>")
}

func main() {
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
