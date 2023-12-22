package config

import (
	"test-gonic/config/app_config"
	"test-gonic/config/db_config"
)

func InitConfig() {

	app_config.InitAppConfig()

	db_config.InitDBConfig()

	// log_config.DefaultLogging()
}
