package userController

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserController struct {
	DB *sql.DB
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	row := controller.DB.QueryRow("INSERT INTO Users (id, name) VALUES (3, $1) RETURNING *", "Michl")

	var (
		id   int
		name string
	)
	row.Scan(&id, &name)

	user := User{id, name}
	payload, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func Init(DB *sql.DB) {
	controller := UserController{DB}
	http.HandleFunc("/createUser", controller.CreateUser)
}
