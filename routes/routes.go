package routes

import (
	"github.com/gin-gonic/gin"
	"makkonen.com/go-mongo/controllers"
)

func InitializeRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.NewBook)

	return router
}
