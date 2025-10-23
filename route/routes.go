package route

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/categories", FindAllCategories)
	server.GET("/categories/:id", FindCategoryByID)
	server.POST("/categories", CreateCategory)
	server.PUT("/categories/:id", UpdateCategoryByID)
	server.DELETE("/categories/:id", DeleteCategoryByID)
}