# {{.ProjectName}}

A Go-Nest project

## Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Set up environment variables (optional):
   ```bash
   export PORT=8080
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=gonest
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

- src/: Source code directory
  - app/: Application setup and configuration
  - common/: Shared utilities and components
  - config/: Configuration management
  - modules/: Feature modules (controllers, services, etc.)

## Adding New Features

Use the gonest CLI to generate new components:

```bash
gonest generate module users
gonest generate controller products
gonest generate service orders
``` 