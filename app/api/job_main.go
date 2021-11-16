package main

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	REPOSITORIES_API_MAX_RESULTS      uint   = 1000
	REPOSITORIES_API_RESULTS_PER_PAGE uint   = 100
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	ISSUES_API_CONCURRENT             uint   = 30
)

func UpdateRepositories() {
	LoadEnv()

	gormDB, sqlDB := GetDBClient()
	defer sqlDB.Close()

	// Fetch and save repositories
	uniqueQuery := [...]string{"stars:30..500", "stars:500..1500", "stars:>1500"}
	for _, eachQuery := range uniqueQuery {
		repositories := FetchRepositories("good-first-issues:>1", eachQuery)
		gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&repositories)
	}

	// Adjust number of repositories
	var (
		repositoryCount    int64
		removeRepositories []Repository
	)
	gormDB.Model(&Repository{}).Count(&repositoryCount)
	removeRepositoryCount := int(repositoryCount) - 3*int(REPOSITORIES_API_MAX_RESULTS)
	if removeRepositoryCount > 0 {
		log.Info(fmt.Sprintf("%d repositories will be removed", removeRepositoryCount))
		gormDB.Model(&Repository{}).Limit(removeRepositoryCount).Find(&removeRepositories)
		gormDB.Unscoped().Delete(&removeRepositories)
	}
}

func UpdateIssues() {
	LoadEnv()

	gormDB, sqlDB := GetDBClient()
	defer sqlDB.Close()

	var (
		repositories []Repository
		wg           sync.WaitGroup
	)
	semaphore := make(chan bool, ISSUES_API_CONCURRENT)
	gormDB.Find(&repositories)

	log.Info("Start fetching issues")
	log.Info(fmt.Sprintf("Number of repositories: %d", len(repositories)))
	for _, repository := range repositories {
		wg.Add(1)
		go func(repository Repository) {
			defer wg.Done()
			semaphore <- true

			issues := FetchIssues(repository.Name)
			repository.Issues = issues
			gormDB.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Save(&repository)

			<-semaphore
		}(repository)
	}
	wg.Wait()
}
