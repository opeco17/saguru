package main

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	REPOSITORIES_API_MAX_RESULTS      uint   = 100
	REPOSITORIES_API_RESULTS_PER_PAGE uint   = 100
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	MINI_BATCH_SIZE                   uint   = 50
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
		mutex        = &sync.Mutex{}
	)
	gormDB.Find(&repositories)

	for i := 0; i < len(repositories)/int(MINI_BATCH_SIZE); i++ {
		// Create mini batch of repository
		lower := i * int(MINI_BATCH_SIZE)
		upper := min((i+1)*int(MINI_BATCH_SIZE), len(repositories)-1)
		miniBatchRepositories := repositories[lower:upper]

		// Fetch issues concurrently for each mini batchn
		modelMiniBatchRepositories := make([]Repository, 0, MINI_BATCH_SIZE)
		for _, repository := range miniBatchRepositories {
			wg.Add(1)
			go func(repository Repository) {
				defer wg.Done()
				issues := FetchIssues(repository.Name)
				repository.Issues = issues

				mutex.Lock()
				modelMiniBatchRepositories = append(modelMiniBatchRepositories, repository)
				mutex.Unlock()

			}(repository)
		}
		wg.Wait()
		gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Save(&modelMiniBatchRepositories)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
