package middleware

import (
	"net/http"
	"path/filepath"
	"test-gonic/utils"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file is required",
		})
		return
	}

	isFileValidatedExtension := utils.FileValidationByExtension(fileHeader, []string{".jpg"})
	if !isFileValidatedExtension {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is not validated",
		})
		return
	}

	isFileValidated := utils.FileValidation(fileHeader, []string{"image/jpg", "image/jpeg"})
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is not validated",
		})
		return
	}

	extensionFile := filepath.Ext(fileHeader.Filename)

	fileName := utils.RandomFileName(extensionFile)

	if saveFile := utils.SaveFile(ctx, fileHeader, fileName); !saveFile {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "file failed upload",
		})
		return
	}

	ctx.Set("filename", fileName)
	ctx.Next()
}
