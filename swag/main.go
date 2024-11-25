package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger" // swagger handler
	_ "go-career/swag/docs"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// HelloWorld godoc
// @Summary say hello
// @Description get a greeting message and the name
// @Tags example
// @Accept json
// @Produce json
// @Param name query string false "The person's name to greet"
// @Success 200 {object} map[string]string
// @Router /hello [get]
func HelloWorld(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + name,
	})
}

func main() {
	e := echo.New()

	// Route => handler
	e.GET("/api/v1/hello", HelloWorld)

	// Serve swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Start(":1323")
}
