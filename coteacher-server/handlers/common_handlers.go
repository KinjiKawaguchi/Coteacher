package handlers

import (
    "github.com/gin-gonic/gin"
)

// CORSMiddleware はCORSリクエストを許可するミドルウェア
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
