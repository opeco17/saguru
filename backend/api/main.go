package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/repositories", getRepositories)
	e.GET("/languages", getLanguages)
	e.GET("/licenses", getLicenses)
	e.GET("/labels", getLabels)
	e.GET("/ordermetrics", getOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
