package restAPI

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oggree/restErrors"
	"github.com/spf13/viper"
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

	env := viper.GetString("env")
	if env == "production" {
		Api.Use(middleware.Recover())
	}

	Api.GET("/health", health)

}

func Start() {
	port := viper.GetString("port")
	if port == "" {
		port = "8080"
	}

	if err := Api.Start(fmt.Sprintf(":%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		Api.Logger.Fatal("error while starting server", err)
	}
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I'm Healthy!")
}

type ResponseModel struct {
	Status bool                  `json:"status"`
	Error  *restErrors.RestError `json:"error"`
	Data   interface{}           `json:"data"`
}

// TODO: change everywhere using this function after middleware
func ResponseSuccessful(payload interface{}) ResponseModel {
	response := ResponseModel{
		Status: true,
		Data:   payload,
	}

	return response
}
