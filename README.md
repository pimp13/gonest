# GoNest CLI

A powerful CLI tool for creating and managing Go projects with Nest.js-like architecture.

## Installation

```bash
go install github.com/pimp13/gonest/cmd/gonest@latest
```

## Features

### Create New Project

Create a new Go-Nest project with a complete structure:

```bash
gonest new my-project
```

This will create a new project with:
- Modular architecture
- PostgreSQL database setup with GORM
- Gin web framework
- Environment-based configuration
- Ready-to-use project structure

### Generate Components

Generate various components for your project:

```bash
# Generate a complete module (includes controller and service)
gonest generate module users
# or short version
gonest g m users

# Generate a controller
gonest generate controller products
# or short version
gonest g c products

# Generate a service
gonest generate service orders
# or short version
gonest g s orders
```

## Project Structure

When you create a new project, it will have the following structure:

```
my-project/
├── app/
│   └── app.go           # Application setup and routing
├── common/
│   └── database/        # Database connection and utilities
├── config/
│   └── config.go        # Configuration management
├── modules/             # Feature modules
├── go.mod              # Go module definition
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Development

### Prerequisites
- Go 1.24 or higher
- PostgreSQL

### Environment Variables

The following environment variables can be configured:

```bash
PORT=8080               # Application port
DB_HOST=localhost      # Database host
DB_PORT=5432          # Database port
DB_USER=postgres      # Database user
DB_PASSWORD=postgres  # Database password
DB_NAME=gonest       # Database name
```

### Getting Started

1. Create a new project:
   ```bash
   gonest new my-project
   cd my-project
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT