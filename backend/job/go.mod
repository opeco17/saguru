module opeco17/gitnavi/job

go 1.16

require (
	github.com/google/go-github/v41 v41.0.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gorm.io/gorm v1.22.3
	opeco17/gitnavi/lib v0.0.1
)

replace opeco17/gitnavi/lib v0.0.1 => ../lib
