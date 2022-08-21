package configuration

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Config for the root command, including flags and types that should be
// available to each subcommand.

type mysqlConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Config struct {
	HTTP struct {
		Address string
	}
	MySQL mysqlConfig
}

// loadEnv variables used on the application configuration,
// and unsets them so that no one can use them directly as globals elsewhere on your code.
func LoadEnv() *Config {
	// Defer clean up calls to just after loadEnv returns.
	defer func() {
		os.Unsetenv("MYSQL_USER")
		os.Unsetenv("MYSQL_PASS")
		os.Unsetenv("MYSQL_HOST")
		os.Unsetenv("MYSQL_PORT")
		os.Unsetenv("MYSQL_NAME")
		os.Unsetenv("HTTP_ADDRESS")
	}()
	return &Config{
		HTTP: struct{ Address string }{
			Address: os.Getenv("HTTP_ADDRESS"),
		},
		MySQL: mysqlConfig{
			User: os.Getenv("MYSQL_USER"),
			Pass: os.Getenv("MYSQL_PASS"),
			Host: os.Getenv("MYSQL_HOST"),
			Port: os.Getenv("MYSQL_PORT"),
			Name: os.Getenv("MYSQL_NAME"),
		},
	}
}
