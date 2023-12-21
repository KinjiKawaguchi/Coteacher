package handlers

import (
	"coteacher/models"
	"coteacher/utils"
	"database/sql"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClassByID(c *gin.Context) {
	var class models.Class

	ID := c.Query("ID")

	query := `SELECT * FROM Classes WHERE ID = ?`
	err := h.DB.QueryRow(query, ID).Scan(&class.ID, &class.Name, &class.TeacherID);
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching class", "error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"class": class})
}

func (h *Handler) GetClassListByTeacherID(c *gin.Context) {
	var classes []models.Class

	TeacherID := c.Query("TeacherID")

	query := `SELECT * FROM Classes WHERE TeacherID = ?`
	rows, err := h.DB.Query(query, TeacherID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching class", "error": err.Error()})
		}
		return
	}

	for rows.Next() {
		var class models.Class
		if err := rows.Scan(&class.ID, &class.Name, &class.TeacherID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
			return
		}
		classes = append(classes, class)
	}
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

func (h *Handler) CreateClass(c *gin.Context) {
	var class models.Class

	Name := c.Query("ClassName")
	TeacherID := c.Query("TeacherID")
	ID :=utils.GenerateUUID(h.DB, "Classes")

	query := `INSERT INTO Classes (ID, Name, TeacherID) VALUES (?, ?, ?)`
	_, err := h.DB.Exec(query, ID, Name, TeacherID)
	if err != nil {
		log.Printf("Failed to create class: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "failed to create class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"class": class})
}

func (h *Handler) GetClass(c *gin.Context) {
	TeacherID := c.Query("TeacherID")

	var classes []models.Class

	query := `SELECT * FROM Classes WHERE TeacherID = ?`
	rows, err := h.DB.Query(query, TeacherID)

	if err != nil {
		log.Printf("Error querying classes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get classes"})
		return
	} else {
		for rows.Next() {
			var class models.Class
			if err := rows.Scan(&class.ID, &class.Name, &class.TeacherID); err != nil {
				log.Printf("Error scanning class: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
				return
			}
			classes = append(classes, class)
		}
	}
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}