package server

import (
	"fmt"
	"log"

	"github.com/golinuxcloudnative/go-url-shortener/config"
	httpHealthz "github.com/golinuxcloudnative/go-url-shortener/internal/api/healthz/delivery/http"
	httpURL "github.com/golinuxcloudnative/go-url-shortener/internal/api/url/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewServer(echo *echo.Echo, cfg *config.Config) *server {
	return &server{echo: echo, cfg: cfg}
}

func (s *server) Run() error {
	//Init handlers
	g := s.echo.Group("")
	api := g.Group("/api/v1")
	//Shortener handlers
	gURL := api.Group("/shortener")
	httpURL.MapURLRoutes(gURL)
	gURL.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip}, user_agent=${user_agent}\n",
	}))

	//Health check handler
	gHealthz := g.Group("/healthz")
	httpHealthz.MapHealthzRoutes(gHealthz)

	// if strings.EqualFold(s.cfg.Server.Mode, "development") {
	// 	fmt.Println(utils.BeautyPrint(s.cfg.Server))
	// 	fmt.Println(utils.BeautyPrint(s.cfg.Bot))
	// 	fmt.Println(s.echo.Logger.Level())
	// }
	port := "5100"
	fmt.Println("port: ", s.cfg.Server.Port)
	if s.cfg.Server.Port != "" {
		port = s.cfg.Server.Port
	}
	if err := s.echo.Start(":" + port); err != nil {
		log.Fatalln("Error starting Server: ", err)
		return err
	}

	return nil

}
