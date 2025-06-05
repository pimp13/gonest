package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templates embed.FS

func newProjectCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "new [project-name]",
		Short:   "Create a new Go-Nest project",
		Long:    `Create a new Go-Nest project with the recommended structure and basic setup`,
		Args:    cobra.ExactArgs(1),
		Example: "gonest new my-awesome-project",
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			createProject(projectName)
		},
	}
}

func createProject(projectName string) {
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		return
	}

	// Project structure
	dirs := []string{
		"src/app",
		"src/common/database",
		"src/common/config",
		"src/app",
	}

	// Create directories
	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create files from templates
	files := map[string]string{
		"main.go": `package main

import (
	"{{.ProjectName}}/src/app"
	"{{.ProjectName}}/src/common/database"
	"{{.ProjectName}}/src/common/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize database connection
	database.Init()

	application := app.NewApp(cfg)

	if err := application.Bootstrap(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}`,

		"go.mod": `module {{.ProjectName}}

go 1.24

require (
	github.com/gin-gonic/gin v1.10.1
	github.com/spf13/cobra v1.9.1
	gorm.io/driver/postgres v1.6.0
	github.com/joho/godotenv v1.5.1
	gorm.io/gorm v1.25.7
)`,

		"src/config/config.go": `package config

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Port string
	Name string
	Url string
}

type Config struct {
	AppConfig *AppConfig
}

func LoadConfig() *Config {
	return &Config{
		AppConfig: &AppConfig{
			Port: getEnv("APP_PORT", "3000"),
			Name: getEnv("APP_NAME", "GoNest"),
			Url: getEnv("APP_URL", "http://localhost:3000"),
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}`,

		"src/app/app.go": `package app

import (
	"{{.ProjectName}}/src/common/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func NewApp(cfg *config.Config) *App {
	return &App{
		config: cfg,
		router: gin.Default(),
	}
}

func (a *App) Bootstrap() error {
	// Setup routes
	a.setupRoutes()

	// Start server
	return a.router.Run(fmt.Sprintf(":%s", a.config.Port))
}

func (a *App) setupRoutes() {
	api := a.router.Group("/api")
	
	// Register module routes here
	// Example: users.NewUserModule().RegisterRoutes(api)
}`,

		"src/common/database/database.go": `package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gonest"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}`,

		".gitignore": `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Environment variables
.env

# IDE specific files
.idea/
.vscode/
*.swp
*.swo

# Dependency directories
vendor/

# Build output
bin/
dist/`,
	}

	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	// Create files from map templates
	for filename, content := range files {
		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			fmt.Printf("Error parsing template for %s: %v\n", filename, err)
			continue
		}

		filePath := filepath.Join(projectName, filename)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}

		if err := tmpl.Execute(file, data); err != nil {
			fmt.Printf("Error writing template to %s: %v\n", filename, err)
			file.Close()
			continue
		}

		file.Close()
	}

	// Create README.md from embedded template
	readmeContent, err := templates.ReadFile("templates/readme.md")
	if err != nil {
		fmt.Printf("Error reading README template: %v\n", err)
		return
	}

	tmpl, err := template.New("README.md").Parse(string(readmeContent))
	if err != nil {
		fmt.Printf("Error parsing README template: %v\n", err)
		return
	}

	readmeFile, err := os.Create(filepath.Join(projectName, "README.md"))
	if err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		return
	}
	defer readmeFile.Close()

	if err := tmpl.Execute(readmeFile, data); err != nil {
		fmt.Printf("Error writing README.md: %v\n", err)
		return
	}

	fmt.Printf("Successfully created new Go-Nest project: %s\n", projectName)
	fmt.Println("\nNext steps:")
	fmt.Println("1. cd", projectName)
	fmt.Println("2. go mod tidy")
	fmt.Println("3. go run main.go")
}
