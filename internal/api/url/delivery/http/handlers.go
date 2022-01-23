package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// func NewUrlHandlers () {

// }

func resolve(c echo.Context) error {
	log.Debug("Level debug")
	log.Info("Level info")
	log.Warn("Level Warn")
	return c.JSON(http.StatusOK, "")
}

func create(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func list(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
