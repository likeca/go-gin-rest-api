package configs

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/notoriouscode97/gin-rest-tutorial/internal/app/rest_api/constants"
	"net/http"
	"os"
)

type Config struct {
	Server   serverConfig
	Database databaseConfig
}

type serverConfig struct {
	Address string
}

type databaseConfig struct {
	Username       string
	Password       string
	Host           string
	DatabaseName   string
	DatabaseDriver string
	DatabaseSource string
}

func NewConfig() *Config {
	err := godotenv.Load("configs/dev.env")

	if err != nil {
		panic("Error loading .env file")
	}

	c := &Config{
		Server: serverConfig{
			Address: GetEnvOrPanic(constants.EnvKeys.ServerAddress),
		},
		Database: databaseConfig{
			Username:       GetEnvOrPanic(constants.EnvKeys.DatabaseUsername),
			Password:       GetEnvOrPanic(constants.EnvKeys.DatabasePassword),
			Host:           GetEnvOrPanic(constants.EnvKeys.DatabaseHost),
			DatabaseName:   GetEnvOrPanic(constants.EnvKeys.DatabaseName),
			DatabaseDriver: GetEnvOrPanic(constants.EnvKeys.DBDriver),
			DatabaseSource: GetEnvOrPanic(constants.EnvKeys.DBSource),
		},
	}

	return c
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}

	return value
}

func (conf *Config) CorsNew() gin.HandlerFunc {
	allowedOrigin := GetEnvOrPanic(constants.EnvKeys.CorsAllowedOrigin)

	return cors.New(cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{constants.Headers.Origin},
		ExposeHeaders:    []string{constants.Headers.ContentLength},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowedOrigin
		},
		MaxAge: constants.MaxAge,
	})
}
