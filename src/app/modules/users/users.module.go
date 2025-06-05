package users

import (
	"github.com/pimp13/gonest/src/common/config"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
	config     *config.Config
	service    *UserService
	controller *UserController
}

func NewUserModule(cfg *config.Config) *UserModule {
	service := NewUserService(cfg)
	controller := NewUserController(service)

	return &UserModule{
		config:     cfg,
		service:    service,
		controller: controller,
	}
}

func (m *UserModule) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")

	group.GET("/", m.controller.GetAllUsers)
	group.GET("/:id", m.controller.GetUserByID)
}
