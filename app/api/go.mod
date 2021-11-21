module opeco17/oss-book/api

go 1.16

require (
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.6.1
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gorm.io/gorm v1.22.3
	opeco17/oss-book/lib v0.0.1
)

replace opeco17/oss-book/lib v0.0.1 => ../lib
