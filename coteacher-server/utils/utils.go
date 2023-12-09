package utils

import (
	"database/sql"

	"github.com/google/uuid"
)

// IsUUIDExists は指定されたテーブルに特定のUUIDが存在するかどうかをチェックします。
// tableName はチェックするテーブル名、id はチェックするUUIDです。
func IsUUIDExists(db *sql.DB, tableName string, id uuid.UUID) bool {
    var tempID uuid.UUID
    query := `SELECT ID FROM ` + tableName + ` WHERE ID = ?`
    err := db.QueryRow(query, id).Scan(&tempID)
    return err == nil
}
