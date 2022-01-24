package http

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golinuxcloudnative/go-url-shortener/domain"
	"github.com/golinuxcloudnative/go-url-shortener/pkg/utils"
	"github.com/labstack/echo/v4"
)

func resolve(svc domain.UrlUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := new(UrlResolveRequest)
		if err := c.Bind(url); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		validate := utils.AnyValidator{Validator: validator.New()}

		if err := validate.Validate(url); err != nil {
			returnErrors := utils.ReturnErros(err, validate)
			return c.JSON(http.StatusBadRequest, returnErrors)
		}
		shortner, err := svc.GetURL(url.URL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err))
		}

		if shortner == nil {
			return c.JSON(http.StatusNotFound, echo.NewHTTPError(http.StatusBadRequest, "short url does not exist"))
		}

		return c.JSON(http.StatusOK, shortner)
	}

}

func create(svc domain.UrlUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := new(UrlCreateRequest)
		if err := c.Bind(url); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		validate := utils.AnyValidator{Validator: validator.New()}

		if err := validate.Validate(url); err != nil {
			returnErrors := utils.ReturnErros(err, validate)
			return c.JSON(http.StatusBadRequest, returnErrors)
		}

		shortner, err := svc.GetURL(url.URL)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err))
		}

		if shortner != nil {
			return c.JSON(http.StatusFound, shortner)
		}

		shortner, err = svc.CreateURL(url.URL)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, shortner)
	}

}

func list(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
