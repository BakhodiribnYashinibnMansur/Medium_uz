package handler

import (
	"mediumuz/package/service"
	"mediumuz/util/logrus"

	_ "mediumuz/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	logrus   *logrus.Logger
}

func NewHandler(services *service.Service, logrus *logrus.Logger) *Handler {
	return &Handler{services: services, logrus: logrus}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)

	}
	api := router.Group("/api", handler.userIdentity)
	{
		account := api.Group("/account")
		{
			account.POST("/upload/image", handler.uploadAccountImage)
			account.GET("/verify", handler.verifyEmail)
			account.GET("/resend", handler.resendCodeToEmail)
			account.GET("/recovery/password", handler.recoveryPassword)
		}
	}
	return router
}
