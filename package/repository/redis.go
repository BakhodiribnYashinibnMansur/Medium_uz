package repository

import (
	"mediumuz/util/logrus"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       string
}

func NewRedisDB(redisConfig *RedisConfig, logrus *logrus.Logger) (*redis.Client, error) {
	redisDB, err := strconv.Atoi(redisConfig.DB)
	logrus.Infof("parsing start DB to int")
	if err != nil {
		logrus.Fatalf("error parsing DB %s", err.Error())
		return nil, err
	}
	logrus.Infof("parsing successfull %v", redisDB)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisDB,
	})
	_, err = client.Ping().Result()
	if err != nil {
		logrus.Fatalf("Ping error %v", err.Error())
		return nil, err
	}
	logrus.Info("successfully REDIS PING")
	return client, nil
}
