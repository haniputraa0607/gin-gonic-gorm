package app_config

import "os"

var PORT = ":8080"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"
var SECRET_KEY = "SECRET_KEY"
var LOG_FILE = "logs/gin.log"

func InitAppConfig() {

	envPort := os.Getenv("APP_PORT")
	if envPort != "" {
		PORT = envPort
	}

	envStaticRouteEnv := os.Getenv("STATIC_ROUTE")
	if envStaticRouteEnv != "" {
		STATIC_ROUTE = envStaticRouteEnv
	}

	envStaticDirEnv := os.Getenv("STATIC_DIR")
	if envStaticDirEnv != "" {
		STATIC_DIR = envStaticDirEnv
	}

	envSecretKey := os.Getenv("SECRET_KEY")
	if envSecretKey != "" {
		SECRET_KEY = envSecretKey
	}

	envLogFile := os.Getenv("LOG_FILE")
	if envLogFile != "" {
		LOG_FILE = envLogFile
	}
}
