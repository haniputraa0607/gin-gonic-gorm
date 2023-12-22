package auth_controller

import (
	"test-gonic/database"
	"test-gonic/model"
	"test-gonic/request"
	"test-gonic/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {

	loginRequest := new(request.LoginRequest)

	if errReq := ctx.ShouldBind(&loginRequest); errReq != nil {

		ctx.AbortWithStatusJSON(400, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	user := new(model.User)
	if errUser := database.DB.Table("users").Where("email = ?", loginRequest.Email).Find(&user).Error; errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "Credential not valid",
		})

		return
	}

	if loginRequest.Password != "12345" {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "Invalid password",
		})

		return
	}

	payloads := jwt.MapClaims{
		"id":    user.ID,
		"tes":   22,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&payloads)

	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed generetad token",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"token":   token,
	})

	return
}
