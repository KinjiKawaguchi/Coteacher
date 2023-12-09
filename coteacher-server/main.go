package main

import (
	"coteacher/handlers"
	"coteacher/routes" // Adjust the import path
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Load `.env` file, open DB connection, etc.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	handler := handlers.NewHandler(db)

	router := gin.Default()
	routes.SetupRouter(router, handler) // Handler インスタンスを渡す

	router.Run()
}
