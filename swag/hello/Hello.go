package hello

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/accounts/{id} [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
