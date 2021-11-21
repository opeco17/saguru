package main

import (
	"fmt"
	"opeco17/oss-book/lib"

	"github.com/labstack/echo/v4"
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

	query := gormDB.Debug().Model(&lib.Issue{})
	query.Joins("inner join labels on labels.issue_id = issues.id")
	if len(input.Labels) > 0 {
		query.Where("labels.name in ?", input.Labels)
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

	query := gormDB.Debug().Model(&lib.Repository{})
	query.Joins("inner join issues on issues.repository_id = repositories.id")
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
		query.Where("issues.id in ?", issueIDs)
	}
	query.Distinct("repositories.id")
	query.Limit(RESULTS_PER_PAGE + 1)
	query.Find(&repositories)

	for _, repository := range repositories {
		repositoryIDs = append(repositoryIDs, repository.ID)
	}
	fmt.Printf("Repository: %d record\n", len(repositories))
	return repositoryIDs
}

func fetchRepositoryEntities(c echo.Context, gormDB *gorm.DB, issueIDs []uint, repositoryIDs []uint) []lib.Repository {
	var repositories []lib.Repository

	query := gormDB.Debug().Model(&repositories)
	if issueIDs == nil {
		query.Preload("Issues")
	} else {
		query.Preload("Issues", "id in ?", issueIDs)
	}
	query.Preload("Issues.Labels")
	query.Preload("Issues.Issuer")
	query.Where("id in ?", repositoryIDs)
	query.Find(&repositories)

	for _, repository := range repositories {
		logrus.Info(fmt.Sprintf("%v: %v", repository.Name, len(repository.Issues)))
	}
	return repositories
}
