package server

import (
	httpHealthz "github.com/golinuxcloudnative/go-url-shortener/internal/api/healthz/delivery/http"
	httpURL "github.com/golinuxcloudnative/go-url-shortener/internal/api/url/delivery/http"
	usecaseHealthz "github.com/golinuxcloudnative/go-url-shortener/usecases/healthz"
	usecaseUrl "github.com/golinuxcloudnative/go-url-shortener/usecases/url"
	"github.com/labstack/echo/v4/middleware"
)

func (s *server) mapHandlers() {
	g := s.echo.Group("")
	api := g.Group("/api/v1")
	//Shortener handlers
	gURL := api.Group("/shortener")
	httpURL.MapURLRoutes(gURL, usecaseUrl.NewService(s.urlRepo))
	gURL.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip}, user_agent=${user_agent}\n",
	}))

	//Health check handler
	gHealthz := g.Group("/healthz")

	httpHealthz.MapHealthzRoutes(gHealthz, s.cfg, s.healthz, usecaseHealthz.NewService(s.healthzRepo))
}
