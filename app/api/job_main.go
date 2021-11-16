package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func UpdateRepositories() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
		os.Exit(1)
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	fmt.Println(s3Bucket)

	// Fetch repositories
	repositories := make([]GitHubRepository, 0, 3*REPOSITORIES_API_MAX_RESULTS)
	repositories = append(repositories, FetchRepositoriesBulk("good-first-issues:>1", "stars:10..500")...)
	repositories = append(repositories, FetchRepositoriesBulk("good-first-issues:>1", "stars:500..1000")...)
	repositories = append(repositories, FetchRepositoriesBulk("good-first-issues:>1", "stars:>1000")...)
	_ = repositories
}
