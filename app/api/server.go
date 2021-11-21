package main

import (
	"opeco17/oss-book/lib"

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

	e.Logger.Fatal(e.Start(":8000"))
}
