package handler

import (
	"fmt"
	"mediumuz/docs"

	"mediumuz/package/service"
	"mediumuz/util/logrus"

	"mediumuz/configs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	logrus   *logrus.Logger
	config   *configs.Configs
}

func NewHandler(services *service.Service, logrus *logrus.Logger, config *configs.Configs) *Handler {
	return &Handler{services: services, logrus: logrus, config: config}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	config := handler.config
	fmt.Println(config)
	docs.SwaggerInfo.Title = config.AppName
	docs.SwaggerInfo.Version = config.Version
	docs.SwaggerInfo.Host = config.ServiceHost + config.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
		// recoveryPassword
		auth.GET("/recovery", handler.recoveryForMessageToEmail)
		auth.GET("/recovery-verify", handler.recoveryCheckEmailCode)
		auth.GET("/recovery-password", handler.recoveryPassword)
	}
	api := router.Group("/api", handler.userIdentity)
	{
		account := api.Group("/account")
		{
			account.GET("/resend", handler.resendCodeToEmail)
			account.GET("/verify", handler.verifyEmail)
			account.PUT("/update", handler.updateAccount)
			account.GET("/get", handler.getUser)
			account.GET("/search", handler.searchUser)
			account.PATCH("/upload-image", handler.uploadAccountImage)
		}
	}
	return router
}
