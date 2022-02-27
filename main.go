package main

import (
	"log"
	"net/http"
	DBConfig "web-go/config"
	userController "web-go/controllers/users"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	DB := DBConfig.ConnectDB()
	userController.Init(DB)

	println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
