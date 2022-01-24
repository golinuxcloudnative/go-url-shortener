package http

import (
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/labstack/echo/v4"
)

func MapURLRoutes(g *echo.Group, service domain.UrlUseCase) {
	//g.GET("/*", )
	g.GET("/", resolve(service))
	g.POST("/", create(service))
	g.GET("/all", list)
}
