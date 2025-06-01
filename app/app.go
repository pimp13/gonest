package app

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pimp13/gonest/config"
	"github.com/pimp13/gonest/modules/users"
)

type App struct {
	config *config.Config
	engine *gin.Engine
}

func NewApp(cfg *config.Config) *App {
	engine := gin.Default()

	app := &App{
		config: cfg,
		engine: engine,
	}

	app.initModules()

	return app
}

func (a *App) initModules() {
	// api prefix for routing
	api := a.engine.Group("/api")

	// initialize modules
	userModule := users.NewUserModule(a.config)
	userModule.RegisterRoutes(api)
}

func (a *App) Bootstrap() error {
	port := strconv.Itoa(a.config.Server.Port)
	return a.engine.Run(":" + port)
}
