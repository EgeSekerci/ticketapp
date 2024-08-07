package db

import (
	"database/sql"
	"fmt"
	"log"
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
