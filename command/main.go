package main

import (
	"mediumuz/configs"
	"mediumuz/package/handler"
	"mediumuz/package/repository"
	"mediumuz/package/service"
	"mediumuz/server"
	"mediumuz/util/logrus"

	_ "github.com/lib/pq"
)

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
	handlers := handler.NewHandler(services, logrus)

	server := new(server.Server)
	err = server.Run(configs.ServerPort, handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

	defer logrus.Infof("run server port:%v", configs.ServerPort)
}
