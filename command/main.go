package main

import (
	_ "github.com/lib/pq"
	"log"
	"mediumuz/configs"
	"mediumuz/package/handler"
	"mediumuz/package/repository"
	"mediumuz/package/service"
	"mediumuz/server"
)

func main() {

	configs, err := configs.InitConfig()

	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     configs.DBHost,
		Port:     configs.DBPort,
		Username: configs.DBUsername,
		DBName:   configs.DBName,
		SSLMode:  configs.DBSSLMode,
		Password: configs.DBPassword,
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(configs.ServerPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
