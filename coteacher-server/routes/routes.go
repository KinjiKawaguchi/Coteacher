package routes

import (
	"coteacher/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, handler *handlers.Handler) {
    // CORSのミドルウェアを追加
    router.Use(handlers.CORSMiddleware())

    // 生徒関連のルート
    router.GET("/Student/CheckAccountExist", handler.CheckAcountExist)
    router.POST("/Student/Create", handler.CreateStudent)
    router.GET("/StudentClass/GetParticipatingClass", handler.GetParticipatingClasses)

}


