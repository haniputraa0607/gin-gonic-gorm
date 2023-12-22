package cors_config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfigManual(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credential", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Header", "Content-Type, Content-length")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()

}

var origins = []string{
	"https://domain.com",
	"https://domain-sub.com",
}

func CorsConfig() gin.HandlerFunc {

	config := cors.DefaultConfig()

	// config.AllowAllOrigins = true
	config.AllowOrigins = origins

	return cors.New(config)

}
