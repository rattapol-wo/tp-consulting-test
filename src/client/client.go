package client

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	dbHost := os.Getenv("SQL_HOST")
	dbPort := os.Getenv("SQL_PORT")
	dbUser := os.Getenv("SQL_USER")
	dbPassword := os.Getenv("SQL_PASSWORD")
	dbName := os.Getenv("SQL_DATABASE")

	// สร้าง DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return db, nil
}
