package constants

import "time"

var EnvKeys = envKeys{
	Env:               "ENV",
	ServerAddress:     "SERVER_ADDRESS",
	DatabaseUsername:  "DATABASE_USERNAME",
	DatabasePassword:  "DATABASE_PASSWORD",
	DatabaseHost:      "DATABASE_HOST",
	DatabaseName:      "DATABASE_NAME",
	CorsAllowedOrigin: "CORS_ALLOWED_ORIGIN",
	DBDriver:          "DB_DRIVER",
	DBSource:          "DB_SOURCE",
}

var Headers = headers{
	Origin:        "Origin",
	ContentLength: "Content-Length",
}

var MaxAge = 12 * time.Hour

type envKeys struct {
	Env               string
	ServerAddress     string
	DatabaseUsername  string
	DatabasePassword  string
	DatabaseHost      string
	DatabaseName      string
	CorsAllowedOrigin string
	DBDriver          string
	DBSource          string
}

type headers struct {
	Origin        string
	ContentLength string
}
