package main

import (
	"fmt"
	"opeco17/oss-book/lib"
	"os"

	"github.com/sirupsen/logrus"
)

func InitDB() error {
	logrus.Info("Start initializing DB.")
	gormDB, sqlDB, err := lib.GetDBClient(
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when initializing DB")
	}
	defer sqlDB.Close()
	err = gormDB.AutoMigrate(&lib.Repository{}, &lib.User{}, &lib.Issue{}, &lib.Label{}, &lib.FrontLanguages{})
	if err != nil {
		return fmt.Errorf("error occurred when migrating database\n%s", err)
	}
	return nil
}
