package handler

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/docs"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/service"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/cors"
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
	// LOCAL
	docs.SwaggerInfo.Host = config.ServiceHost + config.HTTPPort

	// FOR HEROKU
	// docs.SwaggerInfo.Host = config.ServiceHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.New()
	router.Use(cors.CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/test", handler.testHttpsHandler)
	router.Static("/public", "./public/")

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp) //DONE
		auth.POST("/sign-in", handler.signIn) //DONE
		// recoveryPassword
		auth.GET("/recovery", handler.recoveryForMessageToEmail)
		auth.GET("/recovery-verify", handler.recoveryCheckEmailCode)
		auth.GET("/recovery-password", handler.recoveryPassword)
	}

	api := router.Group("/api")
	{
		auth := api.Group("/", handler.userIdentity)
		{
			account := auth.Group("/account")
			{
				account.GET("/sendcode", handler.sendCodeToEmail)          //DONE
				account.GET("/verify", handler.verifyEmail)                //DONE
				account.PUT("/update", handler.updateAccount)              //DONE
				account.GET("/get", handler.getUser)                       //DONE
				account.PATCH("/upload-image", handler.uploadAccountImage) //DONE
				account.GET("/following", handler.followingUser)           //DONE
				account.GET("/follower", handler.followerUser)             //DONE
				account.GET("/get-followings", handler.getFollowings)      //DONE
				account.GET("/get-followers", handler.getFollowers)        //DONE
			}

			post := auth.Group("/post")
			{
				post.POST("/create", handler.createPost)             //DONE
				post.PATCH("/upload-image", handler.uploadImagePost) //DONE
				post.PUT("/update", handler.updatePost)              //DONE
				post.DELETE("/delete", handler.deletePost)           //DONE
				post.GET("/like", handler.likePost)                  //DONE
				post.GET("/view", handler.viewPost)                  //DONE
				post.GET("/rating", handler.ratedPost)               //
				post.POST("/commit", handler.commitPost)             //DONE
			}

		}
		ghost := api.Group("/ghost")
		{
			post := ghost.Group("/post")
			{
				post.GET("/get/:id", handler.getPostID) //DONE
				post.GET("/get-commit", handler.getCommits)
			}

			search := ghost.Group("/")
			{
				search.GET("/search", handler.searchAll) //DONE
			}
		}
	}
	return router
}
