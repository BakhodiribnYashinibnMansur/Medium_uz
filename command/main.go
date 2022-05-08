package main

import (
	"os"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/configs"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/handler"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/service"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/server"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	_ "github.com/lib/pq"
)

// @title MediumuZ API
// @version 1.0
// @description API Server for MediumuZ Application
//@termsOfService https://github.com/BakhodiribnYashinibnMansur/Medium_uz
// @host localhost:8080
// @BasePath /
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	logrus := logrus.GetLogger()
	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	logrus.Info("successfull checked configs.")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		Username: configs.DBUsername,
		DBName:   configs.DBName,
		SSLMode:  configs.DBSSLMode,
		Password: configs.DBPassword,
	}, logrus)

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	logrus.Info("successfull connection DB")
	redis, err := repository.NewRedisDB(&repository.RedisConfig{Host: configs.RedisHost, Port: configs.RedisPort, Password: configs.RedisPassword, DB: configs.RedisDB}, logrus)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db, redis)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logrus, configs)
	port := os.Getenv("APP_PORT")

	server := new(server.Server)
	err = server.Run("8080", handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

	defer logrus.Infof("run server port:%v", port)
}
