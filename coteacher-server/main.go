package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Student struct {
	StudentID int64
	Email     string
	Name      string
}

func main() {
	// Load in the `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// Build router & define routes
	router := gin.Default()
	router.GET("/CheckAcountExist/:email", CheckAcountExist)

	// Run the router
	router.Run()
}

func CheckAcountExist(c *gin.Context) {
	email := c.Param("email")
	email = strings.ReplaceAll(email, "/", "")

	// No need to convert email to an integer.
	var student Student
	query := `SELECT * FROM Student WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&student.StudentID, &student.Email, &student.Name)
	if err != nil {
		c.JSON(404, gin.H{
			"exist": false,
		})
	} else {
		c.JSON(200, gin.H{
			"exist": true,
		})
	}
}
