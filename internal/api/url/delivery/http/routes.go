package http

import "github.com/labstack/echo/v4"

func MapURLRoutes(g *echo.Group) {
	g.GET("/", resolve)
	g.POST("/", create)
	g.GET("/all", list)
}
