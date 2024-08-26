package main

import (
	"SecondHomework/cmd/server"
)

func main() {
	err := server.RunServer()
	if err != nil {
		panic(err)
	}
}
