module opeco17/gitnavi/api

go 1.16

require (
	github.com/labstack/echo/v4 v4.6.1
	github.com/sirupsen/logrus v1.8.1
	gorm.io/gorm v1.22.3
	opeco17/gitnavi/lib v0.0.1
)

replace opeco17/gitnavi/lib v0.0.1 => ../lib
