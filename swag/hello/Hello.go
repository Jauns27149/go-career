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
//
// @Success 201 {object} result.NormalJsonResultWithoutTask{data=response.DedicatedCloud} "Dedicated cloud updated successfully"
//
//	@Router			/accounts/{id} [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// @Tags			goodbye
// @Accept			json
// @Produce		json
//
// @success		200						{object}	result.NormalJsonResult{data=response.DetachVolumeResponse}	"detach volume successfully"
//
// @Param c query echo.Context true "Context"
//
// @Router			/goodbye [get]
func goodbye(c echo.Context) error {
	return c.String(http.StatusOK, "Goodbye, World!")
}
