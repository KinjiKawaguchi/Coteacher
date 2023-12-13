package handlers

import (
	"database/sql"
)

// Handler はハンドラー関数にデータベース接続を提供するための構造体
type Handler struct {
	DB *sql.DB
}

// NewHandler は新しいHandler構造体を作成します。
func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}
