package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
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

	type Data struct {
		name string
		age  int
	}

	data := Data{"John", 30}
	payload, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	// fmt.Fprintf(w, "hello")
}
