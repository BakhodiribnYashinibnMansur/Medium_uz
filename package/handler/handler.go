package handler

import (
	"fmt"
	"mediumuz/package/service"
	"mediumuz/util/logrus"

	"mediumuz/configs"
	_ "mediumuz/docs"

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
	// docs.SwaggerInfo_swagger.Title = config.AppName
	// docs.SwaggerInfo_swagger.Version = config.Version
	// docs.SwaggerInfo_swagger.Host = config.ServiceHost + config.HTTPPort
	// docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}
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
			account.POST("/upload-image", handler.uploadAccountImage)
			account.GET("/verify", handler.verifyEmail)
			account.GET("/resend", handler.resendCodeToEmail)
			account.GET("/recovery-password", handler.recoveryPassword)
			account.POST("/update", handler.updateAccount)
		}
	}
	return router
}
