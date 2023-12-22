package utils

import (
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"test-gonic/config/app_config"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(n int) string {

	rand.Seed(time.Now().UnixMilli())
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {

	contentType := fileHeader.Header.Get("Content-Type")
	// log.Println("content-type", contentType)
	result := false

	for _, typeFile := range fileType {

		if contentType == typeFile {
			result = true
			break
		}
	}

	return result
}

func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {

	extension := filepath.Ext(fileHeader.Filename)
	// log.Println("extension", extension)
	result := false

	for _, typeFile := range fileExtension {

		if extension == typeFile {
			result = true
			break
		}
	}

	return result
}

func RandomFileName(extensionFile string, prefix ...string) string {

	prefixUse := "file"
	if len(prefix) > 0 {
		if prefix[0] != "" {
			prefixUse = prefix[0]
		}
	}

	currentTime := time.Now().UTC().Format("20061206")
	fileName := fmt.Sprintf("%s-%s-%s%s", prefixUse, currentTime, RandomString(5), extensionFile)

	return fileName

}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, fileName string) bool {

	if errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("%s/file/%s", app_config.STATIC_DIR, fileName)); errUpload != nil {
		log.Println("Cant save file")
		return false
	} else {
		return true
	}
}

func RemoveFile(filePath string) error {

	err := os.Remove(filePath)

	if err != nil {
		log.Println("Failed remove file")
		return err
	}

	return nil
}
