package handlers

import (
	"coteacher/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStudentListByClassID(c *gin.Context) {
	classID := c.Query("ClassID")

	var users []models.User

	query := `SELECT * FROM Users WHERE ID IN (SELECT StudentID FROM StudentClasses Where ClassID = ?)`
	rows, err := h.DB.Query(query, classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get participating classes"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error during class iteration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler)GetClassListByStudentID(c *gin.Context) {
	studentID := c.Query("StudentID")

	var classes []models.Class
	query := `SELECT ID, Name, TeacherID FROM Classes WHERE ID IN (SELECT ClassID FROM StudentClasses WHERE StudentID = ?)`
	rows, err := h.DB.Query(query, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get participating classes"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.Name, &class.TeacherID); err != nil {
			log.Printf("Error scanning class: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
			return
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error during class iteration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

func (h *Handler)CreateStudentClass(c *gin.Context) { //授業に参加
	studentID := c.Query("StudentID")
	invitationCode := c.Query("InvitationCode")

	query := `SELECT ID, ClassID, InvitationCode, ExpirationDate, IsActive FROM ClassInvitationCodes WHERE InvitationCode = ? AND ExpirationDate > NOW() AND IsActive = true`
	var classInvitationCode models.ClassInvitationCode
	err := h.DB.QueryRow(query, invitationCode).Scan(&classInvitationCode.ID, &classInvitationCode.ClassID, &classInvitationCode.InvitationCode, &classInvitationCode.ExpirationDate, &classInvitationCode.IsActive)
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

	query = `SELECT 1 FROM StudentClasses WHERE StudentID = ? AND ClassID = ?`
	err = h.DB.QueryRow(query, studentID, classInvitationCode.ClassID).Scan(new(int))
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "student already participated in class"})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error checking class participation"})
		return
	}

	query = `INSERT INTO StudentClasses (StudentID, ClassID) VALUES (?, ?)`
	if _, err := h.DB.Exec(query, studentID, classInvitationCode.ClassID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to participate in class"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "participated in class successfully"})
}

func (h *Handler)DeleteStudentClass(c *gin.Context) { //授業から抜ける

}
