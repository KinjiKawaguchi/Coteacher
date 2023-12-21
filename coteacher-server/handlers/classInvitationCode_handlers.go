package handlers

import (
	"coteacher/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClassInvitationCodeByID(c *gin.Context){
	var classInvitationCode models.ClassInvitationCode
	ID := c.Query("ID")

	query := `SELECT * FROM ClassInvitationCodes WHERE ID = ?`
	err := h.DB.QueryRow(query, ID).Scan(&classInvitationCode.ID, &classInvitationCode.ClassID, &classInvitationCode.InvitationCode, &classInvitationCode.ExpirationDate, &classInvitationCode.IsActive)
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
	c.JSON(http.StatusOK, gin.H{"classInvitationCode": classInvitationCode})
}

func (h *Handler) GetClassInvitationCodeListByClassID(c *gin.Context){
	var classInvitationCodeList []models.ClassInvitationCode
	ClassID := c.Query("ClassID")

	query := `SELECT * FROM ClassInvitationCodes WHERE ClassID = ?`
	rows, err := h.DB.Query(query, ClassID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching class", "error": err.Error()})
		}
		return
	}

	for rows.Next() {
		var classInvitationCode models.ClassInvitationCode
		if err := rows.Scan(&classInvitationCode.ID, &classInvitationCode.ClassID, &classInvitationCode.InvitationCode, &classInvitationCode.ExpirationDate, &classInvitationCode.IsActive); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to scan class"})
			return
		}
		classInvitationCodeList = append(classInvitationCodeList, classInvitationCode)
	}
	c.JSON(http.StatusOK, gin.H{"classInvitationCodes": classInvitationCodeList})
}

func (h *Handler) GetClassInvitationCodeByInvitationCode(c *gin.Context){
	var classInvitationCode models.ClassInvitationCode

	invitationCode := c.Query("InvitationCode")

	query := `SELECT * FROM Classes WHERE InvitationCode = ?`
	err := h.DB.QueryRow(query, invitationCode).Scan(&classInvitationCode.ID, &classInvitationCode.ClassID, &classInvitationCode.InvitationCode, &classInvitationCode.ExpirationDate, &classInvitationCode.IsActive);
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "class not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching class", "error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"classInvitationCode": classInvitationCode})
}

func (h *Handler) CreateClassInvitationCode(c *gin.Context){
	
}

func (h *Handler) UpdateClassInvitationCode(c *gin.Context) {

}