package handler

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/docs"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/service"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/configs"

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
	docs.SwaggerInfo.Title = config.AppName
	docs.SwaggerInfo.Version = config.Version
	docs.SwaggerInfo.Host = config.ServiceHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/test", handler.testHttpsHandler)
	router.Static("/Content", "./public")
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp) //DONE
		auth.POST("/sign-in", handler.signIn) //DONE
		// recoveryPassword
		auth.GET("/recovery", handler.recoveryForMessageToEmail)
		auth.GET("/recovery-verify", handler.recoveryCheckEmailCode)
		auth.GET("/recovery-password", handler.recoveryPassword)
	}
	api := router.Group("/api", handler.userIdentity)
	{
		account := api.Group("/account")
		{
			account.GET("/sendcode", handler.sendCodeToEmail)          //DONE
			account.GET("/verify", handler.verifyEmail)                //DONE
			account.PUT("/update", handler.updateAccount)              //DONE
			account.GET("/get", handler.getUser)                       //DONE
			account.PATCH("/upload-image", handler.uploadAccountImage) //DONE

		}
		post := api.Group("/post")
		{
			post.POST("/create", handler.createPost)             //DONE
			post.GET("/get/:id", handler.getPostID)              //DONE
			post.PATCH("/upload-image", handler.uploadImagePost) //DONE
			post.PUT("/update", handler.updatePost)              //DONE
			post.DELETE("/delete", handler.deletePost)           //DONE
			post.GET("/search", handler.searchAll)               //PROCESS ADVANCED SEARCH
			post.GET("/like")                                    //PROCESS
			post.GET("/view")                                    //PROCESS
		}
	}
	return router
}
