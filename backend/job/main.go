package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "all" {
		initDBAction()
		updateRepositoriesAction()
		updateIssuesAction()
		updateFrontLanguagesAction()
		updateLicensesAction()
		updateLabelsAction()
	} else if flag.Arg(0) == "issue" {
		initDBAction()
		updateIssuesAction()
		updateLabelsAction()
	} else if flag.Arg(0) == "init" {
		initDBAction()
	} else {
		logrus.Error("please specify all, isssue, or init")
		os.Exit(1)
	}
}

func initDBAction() {
	if err := initDB(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateRepositoriesAction() {
	if err := UpdateRepositories(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateIssuesAction() {
	if err := UpdateIssues(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func updateFrontLanguagesAction() {
	if err := UpdateFrontLanguages(); err != nil {
		logrus.Error(err)
	}
}

func updateLicensesAction() {
	if err := UpdateLicenses(); err != nil {
		logrus.Error(err)
	}
}

func updateLabelsAction() {
	if err := UpdateLabels(); err != nil {
		logrus.Error(err)
	}
}
