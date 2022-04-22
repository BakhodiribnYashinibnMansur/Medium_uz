package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	ServerPort    string `default:"8080"`
	DBHost        string `default:"localhost"`
	DBPort        string `default:"2001"`
	DBUsername    string `default:"postgres"`
	DBName        string `default:"mediumuz"`
	DBPassword    string `default:"0224"`
	DBSSLMode     string `default:"false"`
	RedisHost     string `default:"localhost"`
	RedisPort     string `default:"6379"`
	RedisPassword string `default:""`
	RedisDB       string `default:""`
}

func InitConfig() (dbcnfg *Configs, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()

	if err != nil {
		return dbcnfg, fmt.Errorf("fatal error config file: %w ", err)
	}

	if err := godotenv.Load(); err != nil {
		return dbcnfg, fmt.Errorf("error loading env variables: %s", err.Error())
	}

	dbcnfg = &Configs{
		ServerPort:    viper.GetString("port"),
		DBHost:        viper.GetString("db.host"),
		DBPort:        viper.GetString("db.port"),
		DBUsername:    viper.GetString("db.username"),
		DBName:        viper.GetString("db.dbname"),
		DBSSLMode:     viper.GetString("db.sslmode"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		RedisHost:     viper.GetString("redis.host"),
		RedisPort:     viper.GetString("redis.port"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       viper.GetString("redis.db"),
	}
	return
}
