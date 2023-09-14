package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	//MySQL Driver imported - In case you want to use MySQL
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//Postgres Driver imported - In case you want to use Postgres
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

	//	connectionStr := "postgres://postgres:postgres@localhost:5432/go_user_api_db?sslmode=disable"

	d, err := gorm.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}
	db = d

	log.Println("Connected to the database on port 5432 :)")
}

func GetDB() *gorm.DB {
	return db
}
