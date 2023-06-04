package main

import (
	"flag"
	"opeco17/saguru/job/action"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "issue" {
		initMongoDB()
		updateIssues()
		updateCache()
		createIndex()
	} else if flag.Arg(0) == "all" {
		initMongoDB()
		updateRepositories()
		updateIssues()
		updateCache()
		createIndex()
	} else if flag.Arg(0) == "cache" {
		updateCache()
	} else {
		logrus.Error("MustcreateIndex specify all or isssue to a sub command.")
		os.Exit(1)
	}
}

func initMongoDB() {
	if err := action.InitMongoDB(); err != nil {
		logrus.Error("Failed to initialize DB.")
		os.Exit(1)
	}
}

func updateCache() {
	if err := action.UpdateCache(); err != nil {
		logrus.Error("Failed to update caches.")
		os.Exit(1)
	}
}

func createIndex() {
	if err := action.CreateIndex(); err != nil {
		logrus.Error("Failed to create index.")
		os.Exit(1)
	}
}

func updateRepositories() {
	if err := action.UpdateRepositories(); err != nil {
		logrus.Error("Failed to update repositories.")
		os.Exit(1)
	}
}

func updateIssues() {
	if err := action.UpdateIssues(); err != nil {
		logrus.Error("Failed to update issues.")
		os.Exit(1)
	}
}
