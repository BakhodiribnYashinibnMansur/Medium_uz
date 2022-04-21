package handler

import (
	"mediumuz/package/service"
	"mediumuz/util/logrus"

	"github.com/gin-gonic/gin"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	// api := router.Group("/api")
	// {
	// 	lists := api.Group("/lists")
	// 	{
	// 		lists.POST("/")
	// 	}
	// }
	return router
}
