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
	router.Use(CORSMiddleware()) // CORSミドルウェアの追加
	router.GET("/CheckAcountExist/:email", CheckAcountExist)

	// Run the router
	router.Run()
}

// CORSMiddleware はCORSリクエストを許可するミドルウェアです。
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // すべてのオリジンを許可
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func CheckAcountExist(c *gin.Context) {
	email := c.Param("email")
	email = strings.ReplaceAll(email, "/", "")

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

