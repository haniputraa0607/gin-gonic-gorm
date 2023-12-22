package bootstrap

import (
	"log"
	"test-gonic/config"
	"test-gonic/config/app_config"
	"test-gonic/config/cors_config"
	"test-gonic/config/log_config"
	"test-gonic/database"
	"test-gonic/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

	// Get ENV
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get Config
	config.InitConfig()

	/// Connect DB
	database.ConnectDatabase()

	// Create Logging
	log_config.DefaultLogging(app_config.LOG_FILE)

	// Inisiasi GONIC
	app := gin.Default()

	// CORS
	app.Use(cors_config.CorsConfig())

	// Route APPS
	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
