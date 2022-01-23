package server

import (
	"log"

	"github.com/golinuxcloudnative/go-url-shortener/config"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/labstack/echo/v4"
)

type server struct {
	echo        *echo.Echo
	cfg         *config.Config
	healthzRepo domain.HealthzRepository
	urlRepo     domain.UrlRepository
}

func NewServer(echo *echo.Echo, cfg *config.Config, healthzRepo domain.HealthzRepository) *server {
	return &server{echo: echo, cfg: cfg, healthzRepo: healthzRepo}
}

func (s *server) Run() error {
	//Init handlers
	s.mapHandlers()

	// if strings.EqualFold(s.cfg.Server.Mode, "development") {
	// 	fmt.Println(utils.BeautyPrint(s.cfg.Server))
	// 	fmt.Println(utils.BeautyPrint(s.cfg.Bot))
	// 	fmt.Println(s.echo.Logger.Level())
	// }
	port := "5100"
	if s.cfg.Server.Port != "" {
		port = s.cfg.Server.Port
	}
	if err := s.echo.Start(":" + port); err != nil {
		log.Fatalln("Error starting Server: ", err)
		return err
	}

	return nil

}
