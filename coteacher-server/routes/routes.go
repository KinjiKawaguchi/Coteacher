package routes

import (
	"coteacher/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, handler *handlers.Handler) {
	// CORSのミドルウェアを追加
	router.Use(handlers.CORSMiddleware())

	// 生徒関連のルート
	router.GET("/User/Get", handler.GetUser)
	router.POST("/User/Create", handler.CreateUser)
	router.GET("/StudentClass/GetList", handler.GetParticipatingClasses)
	router.GET("/StudentClass/Create", handler.ParticipateClass)

  router.POST("/Class/Create", handler.CreateClass)
	router.GET("/Class/GetList", handler.GetClass)
}
