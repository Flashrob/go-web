package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web-go/config"
)

type Student struct {
	ID int
	Name string
}

func main() {
	DB := DBConfig.ConnectDB()

	rows, err := DB.Query("SELECT * FROM Students")
	if (err != nil) {
		log.Fatal(err);
	}

	students := []Student{}

	for rows.Next() {
		var (
			id int
			name string
		)

		rows.Scan(&id, &name)

		students = append(students, Student{id, name})
	}

	fmt.Println(students)

	http.HandleFunc("/hello", helloHandler)

	println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "hello world yoooo"

	payload, _ := json.Marshal(resp)

	w.Write(payload)
}
