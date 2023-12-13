package handlers

import (
	"coteacher/models"
	"coteacher/utils"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CheckAcountExist(c *gin.Context) {
	email := c.Query("Email")

	var student models.Student
	query := `SELECT * FROM Students WHERE Email = ?`
	err := h.DB.QueryRow(query, email).Scan(&student.ID, &student.Email, &student.Name)
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

func (h *Handler) CreateStudent(c *gin.Context) {
	var student models.Student

	// クエリパラメーターからEmailとNameを取得
	email := c.Query("Email")
	name := c.Query("Name")

	// UUIDを生成
	for {
		student.ID = uuid.New()
		if !utils.IsUUIDExists(h.DB, "Students", student.ID) {
			break
		}
	}

	// データベースに生徒を追加
	query := `INSERT INTO Students (ID, Email, Name) VALUES (?, ?, ?)`
	if _, err := h.DB.Exec(query, student.ID, email, name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student created successfully"})
}

func (h *Handler) GetParticipatingClasses(c *gin.Context) {
	studentID := c.Query("StudentID")

	var classes []models.Class
	query := `SELECT * FROM Classes WHERE ID IN (SELECT ClassID FROM StudentClasses WHERE StudentID = ?)`
	rows, err := h.DB.Query(query, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get participating class"})
		return
	} else {
		// デバッグ出力
		for rows.Next() {
			var class models.Class
			err := rows.Scan(&class.ID, &class.Name, &class.TeacherID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
				return
			}
			classes = append(classes, class)
		}
		c.JSON(http.StatusOK, gin.H{"classes": classes})
	}
}

func (h *Handler) ParticipateClass(c *gin.Context) {
	studentID := c.Query("StudentID")
	invitationCode := c.Query("InvitationCode")
	var classInvitationCode models.ClassInvitationCode

	var dateString string

	query := `SELECT * FROM ClassInvitationCodes WHERE InvitationCode = ? AND ExpirationDate > NOW() AND IsActive = true`
	err := h.DB.QueryRow(query, invitationCode).Scan(&classInvitationCode.ID, &classInvitationCode.ClassID, &classInvitationCode.InvitationCode, &dateString, &classInvitationCode.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Invalid or expired invitation code: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"message": "invalid or expired invitation code"})
		} else {
			log.Printf("Failed to verify invitation code: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to verify invitation code"})
		}
		return
	}
	classInvitationCode.ExpirationDate, err = time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		// Handle parsing error
		log.Printf("Failed to parse expiration date: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to parse expiration date"})
		return
	}

	var exists int
	query = `SELECT 1 FROM StudentClasses WHERE StudentID = ? AND ClassID = ?`
	err = h.DB.QueryRow(query, studentID, classInvitationCode.ClassID).Scan(&exists)
	if err != sql.ErrNoRows {
		c.JSON(http.StatusConflict, gin.H{"message": "student already participated in class"})
		return
	}

	query = `INSERT INTO StudentClasses (StudentID, ClassID) VALUES (?, ?)`
	if _, err := h.DB.Exec(query, studentID, classInvitationCode.ClassID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to participate in class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "participated in class successfully"})
}
