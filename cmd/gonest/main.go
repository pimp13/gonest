package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:     "gonest",
		Version: version,
		Short:   "A CLI tool for Go-Nest framework",
		Long:    `A CLI tool for Go-Nest framework that helps you create modules, controllers, services, and more.`,
	}

	// Add commands
	rootCmd.AddCommand(generateCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
