package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "init" {
		initDBAction()
	} else if flag.Arg(0) == "index" {
		createIndex()
	} else if flag.Arg(0) == "cache" {
		updateCache()
	} else if flag.Arg(0) == "repository" {
		initDBAction()
		updateRepositoriesAction()
		updateCache()
		createIndexAction()
	} else if flag.Arg(0) == "issue" {
		initDBAction()
		updateIssuesAction()
		updateCache()
		createIndexAction()
	} else if flag.Arg(0) == "all" {
		initDBAction()
		updateRepositoriesAction()
		updateIssuesAction()
		updateCache()
		createIndexAction()
	} else {
		logrus.Error("Must specify all, isssue, or init as a sub command.")
		os.Exit(1)
	}
}

func initDBAction() {
	if err := initDB(); err != nil {
		logrus.Error("Failed to initialize DB.")
		os.Exit(1)
	}
}

func updateCacheAction() {
	if err := updateCache(); err != nil {
		logrus.Error("Failed to update caches.")
		os.Exit(1)
	}
}

func createIndexAction() {
	if err := createIndex(); err != nil {
		logrus.Error("Failed to create index.")
		os.Exit(1)
	}
}

func updateRepositoriesAction() {
	if err := updateRepositories(); err != nil {
		logrus.Error("Failed to update repositories.")
		os.Exit(1)
	}
}

func updateIssuesAction() {
	if err := updateIssues(); err != nil {
		logrus.Error("Failed to update issues.")
		os.Exit(1)
	}
}
