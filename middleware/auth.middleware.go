package middleware

import (
	"net/http"
	"strings"
	"test-gonic/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	bearerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	decodeToken, errDecode := utils.DecodeToken(token)
	if errDecode != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	ctx.Set("decode_token", decodeToken)
	ctx.Set("user_id", decodeToken["id"])
	ctx.Set("user_emaiil", decodeToken["emaiil"])
	ctx.Set("user_name", decodeToken["name"])

	ctx.Next()

}

func TokenAuthMiddleware(ctx *gin.Context) {

	token := ctx.GetHeader("X-Token")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	ctx.Next()

}
