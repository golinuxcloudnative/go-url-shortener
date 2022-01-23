package http

import (
	"github.com/golinuxcloudnative/go-url-shortener/config"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/labstack/echo/v4"
)

func MapHealthzRoutes(g *echo.Group, cfg *config.Config, service domain.HealthzUseCase) {
	g.GET("", healthz(cfg, service))
}
