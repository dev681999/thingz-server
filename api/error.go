package main

import "github.com/labstack/echo"

const (
	errBindValue = "invalid value"
)

func (a *app) makeError(c echo.Context, code int, err error) error {
	return c.JSON(code, echo.Map{
		"error": err.Error(),
	})
}
