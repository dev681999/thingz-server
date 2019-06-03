package main

import (
	"github.com/labstack/echo"
)

func (a *app) sendSucess(c echo.Context, data interface{}) error {
	return c.JSON(200, data)
}
