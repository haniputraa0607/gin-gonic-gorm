package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_NAME = "go_gin_gonic"
var DB_USER = "root"
var DB_PASSWORD = ""

func InitDBConfig() {

	envDriver := os.Getenv("DB_DRIVER")
	if envDriver != "" {
		DB_DRIVER = envDriver
	}

	envHost := os.Getenv("DB_HOST")
	if envHost != "" {
		DB_HOST = envHost
	}

	envPort := os.Getenv("DB_PORT")
	if envPort != "" {
		DB_PORT = envPort
	}

	envName := os.Getenv("DB_NAME")
	if envName != "" {
		DB_NAME = envName
	}

	envUser := os.Getenv("DB_USER")
	if envUser != "" {
		DB_USER = envUser
	}

	envPassword := os.Getenv("DB_PASSWORD")
	if envPassword != "" {
		DB_PASSWORD = envPassword
	}

}
