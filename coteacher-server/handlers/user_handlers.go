package handlers

import (
	"coteacher/models"
	"coteacher/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	email := c.Query("Email")

	var user models.User
	query := `SELECT * FROM Users WHERE Email = ?`
	err := h.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.UserType)

	// SQLエラーの適切な処理
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが見つからない場合
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			// その他のデータベースエラー
			log.Printf("Error fetching user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching user"})
		}
		return
	}

	// ユーザーが見つかった場合
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User

	email := c.Query("Email")
	name := c.Query("Name")
	userType := c.Query("UserType")

	// ユーザタイプのバリデーション
	if userType != "Teacher" && userType != "Student" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user type"})
		return
	}

	// ユーザが既に存在するか確認
	query := `SELECT * FROM Users WHERE Email = ?`
	err := h.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.UserType)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
		return
	}

	// errがsql.ErrNoRows以外の場合は、予期せぬエラー
	if err != sql.ErrNoRows {
		log.Printf("Failed to check user existence: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to check user existence"})
		return
	}

	// ユーザーの作成
	uuid := utils.GenerateUUID(h.DB, "Users")
	query = `INSERT INTO Users (ID, Email, Name, UserType) VALUES (?, ?, ?, ?)`
	_, err = h.DB.Exec(query, uuid, email, name, userType)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{"user": map[string]interface{}{
		"ID":       uuid,
		"Email":    email,
		"Name":     name,
		"UserType": userType,
	}})
}
