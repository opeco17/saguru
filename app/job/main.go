package main

import (
	"flag"
	"opeco17/oss-book/lib"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	lib.LoadEnv()

	flag.Parse()
	if flag.Arg(0) == "all" {
		initDB()
		updateRepositories()
		updateIssues()
		updateLanguages()
		updateLabels()
		updateLicenses()
	} else if flag.Arg(0) == "issue" {
		initDB()
		updateIssues()
		updateLanguages()
		updateLabels()
		updateLicenses()
	} else if flag.Arg(0) == "init" {
		InitDB()
	} else {
		logrus.Error("please specify all, isssue, or init")
		os.Exit(1)
	}
}

func initDB() {
	if err := InitDB(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateRepositories() {
	if err := UpdateRepositories(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateIssues() {
	if err := UpdateIssues(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateLanguages() {
	if err := UpdateLanguages(); err != nil {
		logrus.Error(err)
	}
}

func updateLabels() {
	if err := UpdateLabels(); err != nil {
		logrus.Error(err)
	}
}

func updateLicenses() {
	if err := UpdateLicenses(); err != nil {
		logrus.Error(err)
	}
}