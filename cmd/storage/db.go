package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	const targetEnvName = "GO_ENV"
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	filePath := fmt.Sprintf(".env.%s", os.Getenv(targetEnvName))
	fmt.Printf("filePath: %#v\n", filePath)
	// Reads env file and loads them into ENV for this process.
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
