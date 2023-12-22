package file_controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"test-gonic/constanta"
	"test-gonic/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SendStatus(ctx *gin.Context) {

	filename := ctx.MustGet("filename").(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "file is upload",
		"filename": filename,
	})
	return
}

func HandleUploadFile(ctx *gin.Context) {

	decodeToken := ctx.MustGet("decode_token").(jwt.MapClaims)
	fmt.Println("Decode Token => id => ", reflect.ValueOf(decodeToken["id"]).Type())
	fmt.Println("Decode Token => tes => ", reflect.ValueOf(decodeToken["tes"]).Type())
	fmt.Println("Decode Token => email => ", reflect.ValueOf(decodeToken["email"]).Type())
	fmt.Println("Decode Token => name => ", reflect.ValueOf(decodeToken["name"]).Type())

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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file uploaded",
	})
}

func HandleRemoveFile(ctx *gin.Context) {

	filename := ctx.Param("filename")

	if filename == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "filename riquired",
		})
		return
	}

	if err := utils.RemoveFile(constanta.DIR_PUBLIC + constanta.DIR_PUBLIC_FILE + filename); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "file failed delete",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file removed",
	})
}
