module opeco17/oss-book/job

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	gorm.io/gorm v1.22.3
	opeco17/oss-book/lib v0.0.1
)

replace opeco17/oss-book/lib v0.0.1 => ../lib
