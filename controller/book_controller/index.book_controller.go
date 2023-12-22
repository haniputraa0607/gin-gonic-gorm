package book_controller

import "github.com/gin-gonic/gin"

func GetAllBooks(ctx *gin.Context) {

	isValidated := true

	if !isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"hello": "book",
	})
}
