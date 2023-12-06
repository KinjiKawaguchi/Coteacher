package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Students struct {
	ID uuid.UUID
	Email     string
	Name      string
}

type Teachers struct {
	ID uuid.UUID
	Email     string
	Name      string
}

type Classes struct {
	ID uuid.UUID
	Name string
	TeacherID uuid.UUID
}

type StudentClasses struct {
	ClassID uuid.UUID
	StudentID uuid.UUID
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
	router.GET("/Student/CheckAcountExist/:email", checkAcountExist)
  	router.POST("/Student/Create", createStudent) // ルーティングを変更
	router.GET("/StudentClass/GetParticipatingClass", getParticipatingClass)
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

func checkAcountExist(c *gin.Context) {
	email := c.Param("email")
	email = strings.ReplaceAll(email, "/", "")

	var student Students
	query := `SELECT * FROM Student WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&student.ID, &student.Email, &student.Name)
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

func createStudent(c *gin.Context) {
    var student Students

    // クエリパラメーターからEmailとNameを取得
    email := c.Query("Email")
    name := c.Query("Name")

    // UUIDを生成
    for {
        student.ID = uuid.New()
        if !isUUIDExists(student.ID) {
            break
        }
    }

    // データベースに生徒を追加
    query := `INSERT INTO Student (StudentID, Email, Name) VALUES (?, ?, ?)`
    if _, err := db.Exec(query, student.ID, email, name); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create student"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "student created successfully"})
}

func getParticipatingClass(c *gin.Context) {
}

func isUUIDExists(id uuid.UUID) bool {
    var tempID uuid.UUID
    query := `SELECT StudentID FROM Student WHERE StudentID = ?`
    err := db.QueryRow(query, id).Scan(&tempID)
    return err == nil
}
