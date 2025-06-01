package config

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}
type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port: 9000,
		},
		Database: &DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Name:     "go_nest",
		},
	}
}
