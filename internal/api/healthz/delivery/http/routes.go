package http

import "github.com/labstack/echo/v4"

func MapHealthzRoutes(g *echo.Group) {
	g.GET("", healthz)
}
