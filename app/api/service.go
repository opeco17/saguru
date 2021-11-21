package main

import (
	"fmt"
	"opeco17/oss-book/lib"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func fetchIssueIDs(gormDB *gorm.DB, input *getRepositoriesInput) []uint {
	var (
		issues   []lib.Issue
		issueIDs []uint
	)
	if len(input.Labels) == 0 && input.Assigned == nil {
		return nil
	}

	query := gormDB.Model(&lib.Issue{})
	query.Joins("INNER JOIN labels ON labels.issue_id = issues.id")
	if len(input.Labels) > 0 {
		query.Where("labels.name IN ?", input.Labels)
	}
	if *input.Assigned {
		query.Where("issues.assignees_count > ?", 0)
	} else if !*input.Assigned {
		query.Where("issues.assignees_count = ?", 0)
	}
	query.Distinct("issues.id")
	query.Find(&issues)

	for _, issue := range issues {
		issueIDs = append(issueIDs, issue.ID)
	}
	fmt.Printf("Issue: %d record\n", len(issues))
	return issueIDs
}

func fetchRepositoryIDs(gormDB *gorm.DB, input *getRepositoriesInput, issueIDs []uint) []uint {
	var (
		repositories  []lib.Repository
		repositoryIDs []uint
	)

	query := gormDB.Model(&lib.Repository{})
	query.Joins("INNER JOIN issues ON issues.repository_id = repositories.id")
	if input.Language != "" {
		query.Where("repositories.language = ?", input.Language)
	}
	if input.StarCountLower != nil {
		query.Where("repositories.star_count > ?", *input.StarCountLower)
	}
	if input.StarCountUpper != nil {
		query.Where("repositories.star_count < ?", *input.StarCountUpper)
	}
	if input.ForkCountLower != nil {
		query.Where("repositories.fork_count > ?", *input.ForkCountLower)
	}
	if input.ForkCountUpper != nil {
		query.Where("repositories.fork_count < ?", *input.ForkCountUpper)
	}
	if issueIDs != nil {
		query.Where("issues.id IN ?", issueIDs)
	}
	query.Distinct("repositories.id")
	query.Limit(int(RESULTS_PER_PAGE) + 1)
	query.Find(&repositories)

	for _, repository := range repositories {
		repositoryIDs = append(repositoryIDs, repository.ID)
	}
	fmt.Printf("Repository: %d record\n", len(repositories))
	return repositoryIDs
}

func fetchRepositoryEntities(gormDB *gorm.DB, issueIDs []uint, repositoryIDs []uint) []lib.Repository {
	var repositories []lib.Repository

	query := gormDB.Model(&repositories)
	if issueIDs == nil {
		query.Preload("Issues")
	} else {
		query.Preload("Issues", "id IN ?", issueIDs)
	}
	query.Preload("Issues.Labels")
	query.Preload("Issues.Issuer")
	query.Where("id IN ?", repositoryIDs)
	query.Find(&repositories)

	for _, repository := range repositories {
		logrus.Info(fmt.Sprintf("%v: %v", repository.Name, len(repository.Issues)))
	}
	return repositories
}

func fetchFrontLanguages(gormDB *gorm.DB) []lib.FrontLanguages {
	var frontLanguages []lib.FrontLanguages
	gormDB.Model(&frontLanguages).Where("repository_count > ?", MINIMUM_REPOSITORY_COUNT_IN_FRONT_LANGUAGES).Find(&frontLanguages)
	return frontLanguages
}
