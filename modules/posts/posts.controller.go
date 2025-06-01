package posts

import (
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	service *PostsService
}

func NewPostsController(service *PostsService) *PostsController {
	return &PostsController{
		service: service,
	}
}

func (c *PostsController) GetAll(ctx *gin.Context) {
	// TODO: Implement
	ctx.JSON(200, gin.H{"message": "GetAll Posts endpoint"})
}

func (c *PostsController) GetByID(ctx *gin.Context) {
	// TODO: Implement
	ctx.JSON(200, gin.H{"message": "GetByID Posts endpoint"})
}