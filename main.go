package restAPI

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

var Api *echo.Echo

func Init() {
	Api = echo.New()

	Api.GET("/", health)

	Api.GET("/swagger/*", echoSwagger.WrapHandler)

	Api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"*"},
	}))

	Api.Use(middleware.Logger())
	Api.Use(middleware.Recover())

	Api.GET("/health", health)

}

func Start() {
	if err := Api.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		Api.Logger.Fatal("error while starting server", err)
	}
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I'm Healthy!")
}
