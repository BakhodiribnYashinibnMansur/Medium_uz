package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	AppName         string
	ServiceHost     string
	HTTPPort        string
	Version         string
	DBHost          string
	DBPort          string
	DBUsername      string
	DBName          string
	DBPassword      string
	DBSSLMode       string
	RedisHost       string
	RedisPort       string
	RedisPassword   string
	RedisDB         string
	SMTPHost        string
	SMTPPort        int
	SMTPsenderEmail string
	STMPappPassword string
}

func InitConfig() (cfg *Configs, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()

	if err != nil {
		return cfg, fmt.Errorf("fatal error config file: %w ", err)
	}

	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("error loading env variables: %s", err.Error())
	}

	cfg = &Configs{
		AppName:         viper.GetString("app.name"),
		ServiceHost:     viper.GetString("app.server"),
		HTTPPort:        viper.GetString("app.port"),
		Version:         viper.GetString("app.version"),
		DBHost:          viper.GetString("db.host"),
		DBPort:          viper.GetString("db.port"),
		DBUsername:      viper.GetString("db.username"),
		DBName:          viper.GetString("db.dbname"),
		DBSSLMode:       viper.GetString("db.sslmode"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		RedisHost:       viper.GetString("redis.host"),
		RedisPort:       viper.GetString("redis.port"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
		RedisDB:         viper.GetString("redis.db"),
		SMTPHost:        viper.GetString("smtp.host"),
		SMTPPort:        viper.GetInt("smtp.port"),
		SMTPsenderEmail: viper.GetString("smtp.senderemail"),
		STMPappPassword: os.Getenv("SMTP_APP_PASSWORD"),
	}
	return
}
