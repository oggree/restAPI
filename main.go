package restAPI

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

var restAPI *echo.Echo

func Init() {
	restAPI = echo.New()

	restAPI.GET("/", health)

	restAPI.GET("/swagger/*", echoSwagger.WrapHandler)

	restAPI.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"*"},
	}))

	//restAPI.Use(JSONDataMiddleware)

	restAPI.Use(middleware.Logger())
	restAPI.Use(middleware.Recover())

	restAPI.GET("/health", health)

}

func Start() {
	if err := restAPI.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		restAPI.Logger.Fatal("error while starting server", err)
	}
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I'm Healthy!")
}
