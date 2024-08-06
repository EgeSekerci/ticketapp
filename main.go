package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

func loadDbConfig() dbConfig {
	return dbConfig{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbName:     os.Getenv("DB_NAME"),
	}
}

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env")
	}

	configs := loadDbConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.dbHost, configs.dbPort, configs.dbUser, configs.dbPassword, configs.dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error connecting to the database")
	}

	return db
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		db := Connect()
		if db.Ping() != nil {
			fmt.Println("Error connecting database")
		} else {
			fmt.Println("Connection to the database is successful")
		}
	})

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error starting the server")
	}
}
