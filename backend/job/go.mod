module opeco17/saguru/job

go 1.16

require (
	github.com/google/go-github/v41 v41.0.0
	github.com/sirupsen/logrus v1.8.1
	go.mongodb.org/mongo-driver v1.9.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gorm.io/gorm v1.22.3
	opeco17/saguru/lib v0.0.1
)

replace opeco17/saguru/lib v0.0.1 => ../lib
