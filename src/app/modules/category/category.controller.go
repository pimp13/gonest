package category

import (
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *CategoryService
}

func NewCategoryController(service *CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (c *CategoryController) GetAll(ctx *gin.Context) {
	// TODO: Implement
	ctx.JSON(200, gin.H{"message": "GetAll Category endpoint"})
}

func (c *CategoryController) GetByID(ctx *gin.Context) {
	// TODO: Implement
	ctx.JSON(200, gin.H{"message": "GetByID Category endpoint"})
}