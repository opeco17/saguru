package main

import (
	"net/http"
	"opeco17/gitnavi/lib"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	lib.LoadEnv()

	if err := lib.LoadEnv(); err != nil {
		logrus.Fatal("failed to load .env")
	}

	e := echo.New()
	if os.Getenv("LOCAL_RUN") == "true" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{FRONTEND_ORIGIN},
			AllowMethods: []string{http.MethodGet},
		}))
	}
	e.GET("/", index)
	e.GET("/repositories", getRepositories)
	e.GET("/languages", getLanguages)
	e.GET("/licenses", getLicenses)
	e.GET("/labels", getLabels)
	e.GET("/ordermetrics", getOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
