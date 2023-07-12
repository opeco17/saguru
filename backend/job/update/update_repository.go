package update

import (
	"context"
	"fmt"
	"opeco17/saguru/job/constant"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/mongodb"
	"strings"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type queryOption []string

func (queryOpt queryOption) render() string {
	return strings.Join(queryOpt, " ")
}

type gitHubRepository struct {
	github.Repository
}

func (gitHubRepository *gitHubRepository) toMongoDBRepository() *mongodb.Repository {
	name := ""
	if gitHubRepository.FullName != nil {
		name = *gitHubRepository.FullName
	}

	url := ""
	if gitHubRepository.HTMLURL != nil {
		url = *gitHubRepository.HTMLURL
	}

	description := ""
	if gitHubRepository.Description != nil {
		description = *gitHubRepository.Description
	}

	license := ""
	if gitHubRepository.License != nil && gitHubRepository.License.Name != nil {
		license = *gitHubRepository.License.Name
	}

	language := ""
	if gitHubRepository.Language != nil {
		language = *gitHubRepository.Language
	}

	topics := make([]string, 0)
	if gitHubRepository.Topics != nil {
		topics = gitHubRepository.Topics
	}

	repository := &mongodb.Repository{
		RepositoryID:     *gitHubRepository.ID,
		GitHubCreatedAt:  gitHubRepository.CreatedAt.Time,
		GitHubUpdatedAt:  gitHubRepository.UpdatedAt.Time,
		Name:             name,
		URL:              url,
		Description:      description,
		StarCount:        gitHubRepository.StargazersCount,
		ForkCount:        gitHubRepository.ForksCount,
		OpenIssueCount:   gitHubRepository.OpenIssuesCount,
		License:          license,
		Language:         language,
		IssueInitialized: false,
		Topics:           topics,
		Issues:           []*mongodb.Issue{},
	}
	return repository
}

func UpdateRepositories(client *mongo.Client) error {
	logrus.Info("Start updating repositories")

	for _, queryOpt := range getQueryOptions() {
		for _, gitHubRepo := range fetchGitHubRepositories(queryOpt) {
			repo := gitHubRepo.toMongoDBRepository()
			if err := setExistingIssuesToMongoDBRepository(repo, client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to set existing issues to *mongodb.repository: %s", err.Error()))
				continue
			}
			if err := updateMongoDBRepository(repo, client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to upert repository in MongoDB: %s", err.Error()))
				continue
			}
		}
	}

	return nil
}

func getQueryOptions() []queryOption {
	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
	pushedInThreeMonthAgoOption := "pushed:>" + threeMonthAgo
	isGoodFirstIssueOption := "good-first-issues:>0"

	return []queryOption{
		{"stars:>20000", pushedInThreeMonthAgoOption},
		{"stars:15000..20000", pushedInThreeMonthAgoOption},
		{"stars:10000..15000", pushedInThreeMonthAgoOption},
		{"stars:7000..10000", pushedInThreeMonthAgoOption},
		{"stars:6000..7000", pushedInThreeMonthAgoOption},
		{"stars:5000..6000", pushedInThreeMonthAgoOption},
		{"stars:4000..5000", pushedInThreeMonthAgoOption},
		{"stars:3500..4000", pushedInThreeMonthAgoOption},
		{"stars:2500..3000", pushedInThreeMonthAgoOption},
		{"stars:2000..2500", pushedInThreeMonthAgoOption},
		{"stars:1700..2000", pushedInThreeMonthAgoOption},
		{"stars:1500..1700", pushedInThreeMonthAgoOption},
		{"stars:1300..1500", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:1000..1300", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:600..1000", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:400..600", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:300..400", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:200..300", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:100..200", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
		{"stars:30..100", pushedInThreeMonthAgoOption, isGoodFirstIssueOption},
	}
}

func fetchGitHubRepositories(queryOpt queryOption) []gitHubRepository {
	logrus.Info(fmt.Sprintf("Start fetching repositories with query: %s", queryOpt.render()))
	isInitMessagePrinted := false

	gitHubRepos := make([]gitHubRepository, 0, constant.REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < constant.REPOSITORIES_API_MAX_RESULTS/constant.REPOSITORIES_API_RESULTS_PER_PAGE; page++ {
		ctx := context.Background()
		client := util.GetGitHubClient(ctx)
		searchOpt := &github.SearchOptions{
			Sort: "updated",
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: constant.REPOSITORIES_API_RESULTS_PER_PAGE,
			},
		}
		searchResult, resp, err := client.Search.Repositories(ctx, queryOpt.render(), searchOpt)
		if err != nil {
			logrus.Error(err)
			continue
		}
		if resp.StatusCode >= 400 {
			logrus.Error(fmt.Sprintf("Bad response status code %d\n%v", resp.StatusCode, searchResult))
			continue
		}
		if isInitMessagePrinted == false {
			targetRepoCount := constant.REPOSITORIES_API_MAX_RESULTS
			if *searchResult.Total < constant.REPOSITORIES_API_MAX_RESULTS {
				targetRepoCount = *searchResult.Total
			}
			logrus.Info(fmt.Sprintf("%d/%d will be fetched", targetRepoCount, *searchResult.Total))
			isInitMessagePrinted = true
		}

		for _, gitHubRepo := range searchResult.Repositories {
			gitHubRepos = append(gitHubRepos, gitHubRepository{*gitHubRepo})
		}
	}
	return gitHubRepos
}

func setExistingIssuesToMongoDBRepository(repo *mongodb.Repository, client *mongo.Client) error {
	repositoryCollection := client.Database(mongodb.DATABASE_NAME).Collection(mongodb.REPOSITORY_COLLECTION_NAME)

	var existingRepo *mongodb.Repository
	filter := bson.M{"repository_id": repo.RepositoryID}
	if err := repositoryCollection.FindOne(context.Background(), filter).Decode(existingRepo); err != nil {
		return err
	}
	if existingRepo.RepositoryID != 0 {
		repo.Issues = existingRepo.Issues
	}
	return nil
}

func updateMongoDBRepository(repo *mongodb.Repository, client *mongo.Client) error {
	repositoryCollection := client.Database(mongodb.DATABASE_NAME).Collection(mongodb.REPOSITORY_COLLECTION_NAME)

	filter := bson.M{"repository_id": repo.RepositoryID}
	update := bson.M{"$set": repo}
	if _, err := repositoryCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true)); err != nil {
		return err
	}
	return nil
}
