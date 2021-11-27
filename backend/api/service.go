package main

import (
	"fmt"
	"opeco17/oss-book/lib"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func fetchIssueIDs(gormDB *gorm.DB, input *GetRepositoriesInput) []uint {
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
		query.Where("labels.name IN ?", strings.Split(input.Labels, ","))
	}
	if input.Assigned != nil && *input.Assigned {
		query.Where("issues.assignees_count > ?", 0)
	} else if input.Assigned != nil && !*input.Assigned {
		query.Where("issues.assignees_count = ?", 0)
	}
	query.Distinct("issues.id")
	query.Find(&issues)

	for _, issue := range issues {
		issueIDs = append(issueIDs, issue.ID)
	}
	logrus.Info(fmt.Sprintf("Issue: %d record\n", len(issues)))
	return issueIDs
}

func fetchRepositoryIDs(gormDB *gorm.DB, input *GetRepositoriesInput, issueIDs []uint) []uint {
	var (
		repositories  []lib.Repository
		repositoryIDs []uint
	)

	query := gormDB.Model(&lib.Repository{})
	query.Joins("INNER JOIN issues ON issues.repository_id = repositories.id")
	if input.Languages != "" {
		query.Where("repositories.language IN ?", strings.Split(input.Languages, ","))
	}
	if input.License != "" {
		query.Where("repositories.license = ?", input.License)
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
	query.Order("repositories.id")
	query.Distinct("repositories.id")
	query.Offset(int(*input.Page) * int(RESULTS_PER_PAGE))
	query.Limit(int(RESULTS_PER_PAGE) + 1)
	query.Find(&repositories)

	for _, repository := range repositories {
		repositoryIDs = append(repositoryIDs, repository.ID)
	}
	logrus.Info(fmt.Sprintf("Repository: %d record\n", len(repositories)))
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

func fetchFrontLanguages(gormDB *gorm.DB) []lib.FrontLanguage {
	var frontLanguages []lib.FrontLanguage
	gormDB.Model(&frontLanguages).Where("repository_count > ?", MINIMUM_REPOSITORY_COUNT_IN_FRONT_LANGUAGES).Find(&frontLanguages)
	return frontLanguages
}

func fetchFrontLicenses(gormDB *gorm.DB) []lib.FrontLicense {
	var frontLicenses []lib.FrontLicense
	gormDB.Model(&frontLicenses).Where("repository_count > ?", MINIMUM_REPOSITORY_COUNT_IN_FRONT_LICENSES).Find(&frontLicenses)
	return frontLicenses
}

func fetchFrontLabels(gormDB *gorm.DB) []lib.FrontLabel {
	var frontLabels []lib.FrontLabel
	gormDB.Model(&frontLabels).Where("issue_count > ?", MINIMUM_ISSUE_COUNT_IN_FRONT_LABELS).Find(&frontLabels)
	return frontLabels
}
