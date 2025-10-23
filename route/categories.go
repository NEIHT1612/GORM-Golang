package route

import (
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/repo"
	"example.com/m/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindAllCategories(ctx *gin.Context) {
	repo := repo.NewCategoryRepo()
	service := service.NewCategoryService(repo)
	categories, err := service.FindAllCategories()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data categories"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func FindCategoryByID(ctx *gin.Context){
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}

	repo := repo.NewCategoryRepo()
	service := service.NewCategoryService(repo)
	category, err := service.FindCategoryByID(categoryID)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't get data category"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func CreateCategory(ctx *gin.Context){
	var category models.Category
	result := ctx.ShouldBindJSON(&category)
	if result != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't binding data"})
		return
	}
	
	repo := repo.NewCategoryRepo()
	service := service.NewCategoryService(repo)
	err := service.CreateCategory(category)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't create"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Create successful"})
}

func UpdateCategoryByID(ctx *gin.Context){
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var category models.Category
	result := ctx.ShouldBindJSON(&category)
	if result != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't binding data"})
		return
	}

	repo := repo.NewCategoryRepo()
	service := service.NewCategoryService(repo)

	err = service.UpdateCategoryByID(categoryID, category)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Update failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Update successful"})
}

func DeleteCategoryByID(ctx *gin.Context){
	categoryID, err := uuid.Parse(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return 
	}

	repo := repo.NewCategoryRepo()
	service := service.NewCategoryService(repo)
	err = service.DeleteCategoryByID(categoryID)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Delete failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Successful"})
}