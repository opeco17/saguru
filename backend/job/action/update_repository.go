package action

import (
	"context"
	"fmt"
	"opeco17/saguru/job/constant"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/model"
	"strings"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateRepositories() error {
	logrus.Info("Start updating repositories.")

	// Connect DB
	client, err := util.GetMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	repositoryCollection := client.Database("main").Collection("repositories")

	// Fetch and save repositories
	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
	pushedBeforeThreeMonthAgoQuery := "pushed:>" + threeMonthAgo
	isGoodFirstIssueQuery := "good-first-issues:>0"
	queries := [][]string{
		{"stars:30..100", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:100..200", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:200..300", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:300..400", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:400..600", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:600..1000", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1000..1300", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1300..1500", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1500..1700", pushedBeforeThreeMonthAgoQuery},
		{"stars:1700..2000", pushedBeforeThreeMonthAgoQuery},
		{"stars:2000..2500", pushedBeforeThreeMonthAgoQuery},
		{"stars:2500..3000", pushedBeforeThreeMonthAgoQuery},
		{"stars:3500..4000", pushedBeforeThreeMonthAgoQuery},
		{"stars:4000..5000", pushedBeforeThreeMonthAgoQuery},
		{"stars:5000..6000", pushedBeforeThreeMonthAgoQuery},
		{"stars:6000..7000", pushedBeforeThreeMonthAgoQuery},
		{"stars:7000..10000", pushedBeforeThreeMonthAgoQuery},
		{"stars:10000..15000", pushedBeforeThreeMonthAgoQuery},
		{"stars:15000..20000", pushedBeforeThreeMonthAgoQuery},
		{"stars:>20000", pushedBeforeThreeMonthAgoQuery},
	}
	for _, eachQuery := range queries {
		now := time.Now()
		repositories := fetchRepositories(
			eachQuery...,
		)
		for _, repository := range repositories {
			repository.UpdatedAt = time.Now()

			// Set existing issues
			var oldRepository model.Repository
			findFilter := bson.M{"repository_id": repository.RepositoryID}
			repositoryCollection.FindOne(context.TODO(), findFilter).Decode(&oldRepository)
			if oldRepository.RepositoryID != 0 {
				repository.Issues = oldRepository.Issues
			}

			// Update
			updateFilter := bson.M{"repository_id": repository.RepositoryID}
			update := bson.M{"$set": repository}
			_, err = repositoryCollection.UpdateOne(context.TODO(), updateFilter, update, options.Update().SetUpsert(true))
			if err != nil {
				logrus.Warn(err)
			}
		}
		if restTimeSecond := constant.REPOSITORIES_API_INTERVAL_SECOND - int(time.Since(now).Seconds()); restTimeSecond > 0 {
			time.Sleep(time.Second * time.Duration(restTimeSecond))
		}
	}

	// Adjust number of repositories by removing old repositories
	repositoryCount, err := repositoryCollection.CountDocuments(context.TODO(), bson.M{})
	if removeRepositoryCount := int(repositoryCount) - constant.MAX_REPOSITORY_RECORES; removeRepositoryCount > 0 {
		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))

		opts := options.Find().SetLimit(int64(removeRepositoryCount)).SetProjection(bson.M{"_id": 1}).SetSort(bson.M{"git_hub_updated_at": 1})
		cursor, err := repositoryCollection.Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			logrus.Error(err)
			return err
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var result model.Repository
			if err := cursor.Decode(&result); err != nil {
				logrus.Error(err)
				return err
			}
			repositoryCollection.DeleteOne(context.TODO(), result)
		}
	}

	logrus.Info("Successfully finished to update repositories.")
	return nil
}

func fetchRepositories(queries ...string) []*model.Repository {
	gitHubRepositories := fetchGitHubRepositories(queries...)
	repositories := make([]*model.Repository, 0, len(gitHubRepositories))
	for _, gitHubRepository := range gitHubRepositories {
		repositories = append(repositories, convertRepository(gitHubRepository))
	}
	return repositories
}

func fetchGitHubRepositories(queries ...string) []*github.Repository {
	gitHubRepositories := make([]*github.Repository, 0, constant.REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < constant.REPOSITORIES_API_MAX_RESULTS/constant.REPOSITORIES_API_RESULTS_PER_PAGE; page++ {
		gitHubRepositoriesResponse, queries, err := fetchGitHubRepositoriesSubset(page, queries...)
		if err != nil {
			logrus.Error("Failed to fetch repositories from GitHub API")
			continue
		}
		gitHubRepositories = append(gitHubRepositories, gitHubRepositoriesResponse.Repositories...)
		if page == 0 {
			logrus.Info("Start fetching repositories.")
			logrus.Info(fmt.Sprintf("Query: %v", queries))
			logrus.Info(fmt.Sprintf("Total count: %v", *gitHubRepositoriesResponse.Total))
		}
	}
	return gitHubRepositories
}

func fetchGitHubRepositoriesSubset(page int, queries ...string) (*github.RepositoriesSearchResult, string, error) {
	ctx := context.Background()
	client := util.GetGitHubClient(ctx)
	opts := &github.SearchOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: constant.REPOSITORIES_API_RESULTS_PER_PAGE,
		},
	}
	body, resp, err := client.Search.Repositories(ctx, strings.Join(queries, " "), opts)
	if err != nil {
		logrus.Error(err)
		return nil, "", err
	}
	if resp.StatusCode >= 400 {
		message := fmt.Sprintf("bad response status code %d\n%v", resp.StatusCode, body)
		logrus.Error(message)
		return nil, "", fmt.Errorf(message)
	}
	return body, strings.Join(queries, " "), nil
}

func convertRepository(gitHubRepository *github.Repository) *model.Repository {
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

	repository := &model.Repository{
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
		Issues:           []*model.Issue{},
	}
	return repository
}

func convertUser(gitHubUser *github.User) *model.User {
	name := ""
	if gitHubUser.Name != nil {
		name = *gitHubUser.Name
	}

	url := ""
	if gitHubUser.HTMLURL != nil {
		url = *gitHubUser.HTMLURL
	}

	avatarURL := ""
	if gitHubUser.AvatarURL != nil {
		avatarURL = *gitHubUser.AvatarURL
	}

	user := &model.User{
		UserID:    *gitHubUser.ID,
		Name:      name,
		URL:       url,
		AvatarURL: avatarURL,
	}
	return user
}

func convertLabels(gitHubLabels []*github.Label) []*model.Label {
	labels := make([]*model.Label, 0, len(gitHubLabels))
	for _, gitHubLabel := range gitHubLabels {
		name := ""
		if gitHubLabel.Name != nil {
			name = *gitHubLabel.Name
		}
		label := &model.Label{
			LabelID: *gitHubLabel.ID,
			Name:    name,
		}
		labels = append(labels, label)
	}
	return labels
}
