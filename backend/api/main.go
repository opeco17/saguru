package main

import (
	"opeco17/gitnavi/lib"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	lib.LoadEnv()

	if err := lib.LoadEnv(); err != nil {
		logrus.Fatal("failed to load .env")
	}

	e := echo.New()
	e.GET("/", index)
	e.GET("/repositories", getRepositories)
	e.GET("/languages", getLanguages)
	e.GET("/licenses", getLicenses)
	e.GET("/labels", getLabels)
	e.GET("/ordermetrics", getOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
