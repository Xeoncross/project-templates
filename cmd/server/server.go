package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/xeoncross/project-templates/database/db"
	"github.com/xeoncross/project-templates/internal/configuration"
	"github.com/xeoncross/project-templates/internal/httpserver"
	"github.com/xeoncross/project-templates/internal/mysql"
	"github.com/xeoncross/project-templates/internal/service"
)

func run() error {

	config := configuration.LoadEnv()

	{
		enc := json.NewEncoder(os.Stderr)
		enc.SetIndent("", "  ")
		enc.Encode(config)
	}

	my := config.MySQL
	connection, err := mysql.Load(my.User, my.Pass, my.Host, my.Port, my.Name)
	if err != nil {
		return fmt.Errorf("%s: %w", my.Host, err)
	}

	db := db.New(connection)

	handler := httpserver.Handler{S: &service.User{DB: db}}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:email", handler.GetUser)
	e.POST("/users", handler.CreateUser)

	// Limit by api key
	// e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	// 	KeyLookup: "header:api-key",
	// 	Validator: func(key string, c echo.Context) (bool, error) {
	// 		return key == "api-key", nil
	// 	},
	// }))

	return e.Start(config.HTTP.Address)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
