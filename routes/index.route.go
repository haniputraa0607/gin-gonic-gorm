package routes

import (
	"test-gonic/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	// Route Static
	app.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	route := app.Group("/api")

	v1Route(route)

}
