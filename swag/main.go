package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-career/swag/docs"
	"net/http"
)

// @title Hello World
// @host localhost:1323
// @BasePath /

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/hello", HelloWorld)
	e.GET("/goodbye", goodbye)

	e.Logger.Fatal(e.Start(":1323"))
}

// HelloWorld
// @Tags Hello
// @Produce json
// @Router /hello [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// @Tags Goodbye
// @Router /goodbye [get]
func goodbye(c echo.Context) error {
	return c.String(http.StatusOK, "Goodbye, World!")
}
