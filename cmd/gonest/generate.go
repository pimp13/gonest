package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate [type] [name]",
		Aliases: []string{"g"},
		Short:   "Generate Go-Nest components",
		Long:    `Generate various components like controllers, services, modules, etc.`,
	}

	cmd.AddCommand(
		generateController(),
		generateService(),
		generateModule(),
		generateRepository(),
	)

	return cmd
}

func generateController() *cobra.Command {
	return &cobra.Command{
		Use:     "controller [name]",
		Aliases: []string{"c"},
		Short:   "Generate a controller",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			generateComponent("controller", name)
		},
	}
}

func generateService() *cobra.Command {
	return &cobra.Command{
		Use:     "service [name]",
		Aliases: []string{"s"},
		Short:   "Generate a service",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			generateComponent("service", name)
		},
	}
}

func generateModule() *cobra.Command {
	return &cobra.Command{
		Use:     "module [name]",
		Aliases: []string{"m"},
		Short:   "Generate a module",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			generateComponent("module", name)
		},
	}
}

func generateRepository() *cobra.Command {
	return &cobra.Command{
		Use:     "repository [name]",
		Aliases: []string{"r"},
		Short:   "Generate a repository",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			generateComponent("repository", name)
		},
	}
}

func generateComponent(componentType, name string) {
	templates := map[string]string{
		"controller": `package {{.Package}}

		import (
			"github.com/gin-gonic/gin"
		)

		type {{.Name}}Controller struct {
			service *{{.Name}}Service
		}

		func New{{.Name}}Controller(service *{{.Name}}Service) *{{.Name}}Controller {
			return &{{.Name}}Controller{
				service: service,
			}
		}

		func (c *{{.Name}}Controller) GetAll(ctx *gin.Context) {
			// TODO: Implement
			ctx.JSON(200, gin.H{"message": "GetAll {{.Name}} endpoint"})
		}

		func (c *{{.Name}}Controller) GetByID(ctx *gin.Context) {
			// TODO: Implement
			ctx.JSON(200, gin.H{"message": "GetByID {{.Name}} endpoint"})
		}`,

		"service": `package {{.Package}}

		type {{.Name}}Service struct {
		}

		func New{{.Name}}Service() *{{.Name}}Service {
			return &{{.Name}}Service{}
		}

		func (s *{{.Name}}Service) GetAll() ([]interface{}, error) {
			// TODO: Implement
			return nil, nil
		}

		func (s *{{.Name}}Service) GetByID(id string) (interface{}, error) {
			// TODO: Implement
			return nil, nil
		}`,

		"module": `package {{.Package}}

		import (
			"github.com/gin-gonic/gin"
		)

		type {{.Name}}Module struct {
			controller *{{.Name}}Controller
			service    *{{.Name}}Service
		}

		func New{{.Name}}Module() *{{.Name}}Module {
			service := New{{.Name}}Service()
			controller := New{{.Name}}Controller(service)

			return &{{.Name}}Module{
				service:    service,
				controller: controller,
			}
		}

		func (m *{{.Name}}Module) RegisterRoutes(router *gin.RouterGroup) {
			group := router.Group("/{{.LowerName}}")

			group.GET("/", m.controller.GetAll)
			group.GET("/:id", m.controller.GetByID)
		}`,

		"repository": `package {{.Package}}

		import (
			"github.com/gin-gonic/gin"
		)

		type {{.Name}}Repository struct {
			// inject database
		}

		func New{{.Name}}Repository() *{{.Name}}Repository {
			return &{{.Name}}Repository{}
		}`,
	}

	// Create module directory
	modulePath := filepath.Join("src/modules", strings.ToLower(name))
	err := os.MkdirAll(modulePath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Prepare template data
	caser := cases.Title(language.English)
	data := struct {
		Name      string
		Package   string
		LowerName string
	}{
		Name:      caser.String(name),
		Package:   strings.ToLower(name),
		LowerName: strings.ToLower(name),
	}		

	// Generate the component
	filename := fmt.Sprintf("%s.%s.go", strings.ToLower(name), componentType)
	filePath := filepath.Join(modulePath, filename)

	tmpl, err := template.New(componentType).Parse(templates[componentType])
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated %s at %s\n", componentType, filePath)
}
