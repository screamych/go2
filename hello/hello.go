package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, Linux!")

	client := &http.Client{Timeout: time.Second}
	resOne, err := client.Get("https://golang.org/doc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resOne.Status, resOne.StatusCode, resOne.Request.URL)
	fmt.Println(resOne.Request.Header, resOne.Request.Response.Header)

	body, _ := io.ReadAll(resOne.Body)
	resOne.Body.Close()

	file, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
