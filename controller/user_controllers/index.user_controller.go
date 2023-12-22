package user_controllers

import (
	"net/http"
	"strconv"
	"test-gonic/database"
	"test-gonic/model"
	"test-gonic/request"
	"test-gonic/response"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {

	users := new([]model.User)

	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetById(ctx *gin.Context) {

	user := new(response.UserResponse)

	id := ctx.Param("id")

	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func Store(ctx *gin.Context) {

	userReq := new(request.UserRequest)

	if err := ctx.ShouldBind(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userEmailExist := new(model.User)
	if database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist); userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
		})
		return
	}

	user := new(model.User)
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	if errDB := database.DB.Table("users").Create(&user).Error; errDB != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errDB.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})

}

func UpdateById(ctx *gin.Context) {

	id := ctx.Param("id")
	user := new(model.User)
	userReq := new(request.UserRequest)
	userEmailExist := new(model.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	if errUser := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; errUser != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	if errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error; errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already exist",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	if errUpdateData := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error; errUpdateData != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdateData.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})

}

func DeleteById(ctx *gin.Context) {

	id := ctx.Param("id")
	user := new(model.User)

	if errUser := database.DB.Table("users").Where("id = ?", id).Find(&user).Error; errUser != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	if errDB := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&model.User{}).Error; errDB != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errDB.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetAllUsersPagination(ctx *gin.Context) {

	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	limit := ctx.Query("limit")
	if limit == "" {
		limit = "10"
	}
	limitInt, _ := strconv.Atoi(limit)

	users := new([]model.User)

	err := database.DB.Table("users").Offset((pageInt - 1) * limitInt).Limit(limitInt).Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": len(*users),
	})
}
