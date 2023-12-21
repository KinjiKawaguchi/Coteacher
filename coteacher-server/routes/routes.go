package routes

import (
	"coteacher/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, handler *handlers.Handler) {
	// CORSのミドルウェアを追加
	router.Use(handlers.CORSMiddleware())

	// 生徒関連のルート
	router.GET("/class/id", handler.GetClassByID)
	router.GET("/class/teacherID", handler.GetClassListByTeacherID)
	router.POST("/class", handler.CreateClass)

	router.GET("/classInvitationCode/id", handler.GetClassInvitationCodeByID)
	router.GET("/classInvitationCode/classID", handler.GetClassInvitationCodeListByClassID)
	router.GET("/classInvitationCode/invitationCode", handler.GetClassInvitationCodeByInvitationCode)
	router.POST("/classInvitationCode", handler.CreateClassInvitationCode)

	router.GET("/studentClass/classID", handler.GetStudentListByClassID)
	router.GET("/studentClass/studentID", handler.GetClassListByStudentID)
	router.POST("/studentClass", handler.CreateStudentClass)
	router.DELETE("/studentClass", handler.DeleteStudentClass)

	router.GET("user/id", handler.GetUserByID)
	router.GET("user/email", handler.GetUserByEmail)
	router.POST("user", handler.CreateUser)
}
