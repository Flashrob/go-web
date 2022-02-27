package DBConfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	dbname = "testgo"
)

var Database *sql.DB

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Cannot connect : %s", err)
	}

	Database := DB
	return Database
}

func GetDatabase() *sql.DB {
	return Database
}
