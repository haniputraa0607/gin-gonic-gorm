package routes

import (
	"test-gonic/controller/auth_controller"
	"test-gonic/controller/book_controller"
	"test-gonic/controller/file_controller"
	"test-gonic/controller/user_controllers"
	"test-gonic/middleware"

	"github.com/gin-gonic/gin"
)

func v1Route(app *gin.RouterGroup) {

	route := app

	//Route Auth
	route.POST("login", auth_controller.Login)

	// Route User
	routeUser := route.Group("user", middleware.AuthMiddleware)
	routeUser.GET("/", user_controllers.GetAllUsers)
	routeUser.GET("/paginate", user_controllers.GetAllUsersPagination)
	routeUser.POST("", user_controllers.Store)
	routeUser.GET("/:id", user_controllers.GetById)
	routeUser.PATCH("/:id", user_controllers.UpdateById)
	routeUser.DELETE("/:id", user_controllers.DeleteById)

	// Route Book
	routeBook := route.Group("book", middleware.AuthMiddleware)
	routeBook.GET("/", book_controller.GetAllBooks)

	// Route File
	routeFile := route.Group("file", middleware.AuthMiddleware)
	routeFile.POST("/", file_controller.HandleUploadFile)
	routeFile.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
	routeFile.DELETE("/:filename", file_controller.HandleRemoveFile)

}
