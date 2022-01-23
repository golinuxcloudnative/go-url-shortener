package http

import (
	"log"
	"net/http"

	"github.com/golinuxcloudnative/go-url-shortener/config"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/labstack/echo/v4"
)

// type healthzHandlers struct {
// 	db     *domain.UrlRepository
// 	cfg    *config.Config
// 	status interface{}
// }

// func NewHealthzHandlers(db *domain.UrlRepository, cfg *config.Config) *healthzHandlers {
// 	return &healthzHandlers{db: db, cfg: cfg}
// }

func healthz(cfg *config.Config, healthz domain.Healthz, svc domain.HealthzUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := svc.Ping()
		if err != nil {
			log.Printf("could not ping database: %v", err)
			healthz.Database = err.Error()
			healthz.Status = domain.Unhealthy
			return c.JSON(http.StatusServiceUnavailable, healthz)
		}

		healthz.Status = domain.Healthy
		healthz.Database = result

		return c.JSON(http.StatusOK, healthz)
	}
}
