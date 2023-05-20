package action

import (
	"context"
	"fmt"
	"opeco17/saguru/job/constant"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/model"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateIssues() error {
	logrus.Info("Start updating issues.")

	// Connect DB
	client, err := util.GetMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	repositoryCollection := client.Database("main").Collection("repositories")

	// Update issues
	for i := 0; i < constant.UPDATE_ISSUE_BATCH_SIZE/constant.UPDATE_ISSUE_MINIBATCH_SIZE; i++ {
		if i%10 == 0 {
			logrus.Info(fmt.Sprintf("Updating issues: %d/%d", i*constant.UPDATE_ISSUE_MINIBATCH_SIZE, constant.UPDATE_ISSUE_BATCH_SIZE))
		}
		if err := updateIssuesMinibach(repositoryCollection, constant.UPDATE_ISSUE_MINIBATCH_SIZE); err != nil {
			return err
		}
	}

	logrus.Info("Successfully finished to update issues.")
	return nil
}

func updateIssuesMinibach(repositoryCollection *mongo.Collection, updateCount int) error {
	type SimpleRepository struct {
		RepositoryID int    `bson:"repository_id,omitempty"`
		Name         string `bson:"name,omitempty"`
	}
	type RepositoryDiff struct {
		UpdatedAt        time.Time      `bson:"updated_at"`
		IssueInitialized bool           `bson:"issue_initialized"`
		Issues           []*model.Issue `bson:"issues"`
	}
	simpleRepoFields := bson.M{"repository_id": 1, "name": 1}

	// Fetch not initialized repositories
	var notInitializedRepositories []*SimpleRepository
	opts := options.Find().SetLimit(int64(updateCount)).SetSort(bson.M{"updated_at": 1}).SetProjection(simpleRepoFields)
	filter := bson.M{"issue_initialized": false}
	initializedIssueCursor, err := repositoryCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if err = initializedIssueCursor.All(context.TODO(), &notInitializedRepositories); err != nil {
		logrus.Error(err)
		return err
	}

	// Fetch initialized repositories
	var initializedRepositories []*SimpleRepository
	if restRepositoryCount := updateCount - len(notInitializedRepositories); restRepositoryCount > 0 {
		opts = options.Find().SetLimit(int64(updateCount)).SetSort(bson.M{"updated_at": 1}).SetProjection(simpleRepoFields)
		filter = bson.M{"issue_initialized": true}
		notInitializedIssueCursor, err := repositoryCollection.Find(context.TODO(), filter, opts)
		if err != nil {
			logrus.Error(err)
			return err
		}
		if err = notInitializedIssueCursor.All(context.TODO(), &initializedRepositories); err != nil {
			logrus.Error(err)
			return err
		}
	}

	// Concat repositories
	repositories := append(notInitializedRepositories, initializedRepositories...)

	// Update issues concurrently
	var wg sync.WaitGroup
	concurrencyLimitCh := make(chan struct{}, constant.FETCH_ISSUE_CONCURRENCY)
	wg.Add(len(repositories))

	// Fetch and update issues
	for _, repository := range repositories {
		go func(repository *SimpleRepository) {
			concurrencyLimitCh <- struct{}{}
			defer wg.Done()
			defer func() { <-concurrencyLimitCh }()

			issues := fetchIssues(repository.Name)
			diff := RepositoryDiff{
				UpdatedAt:        time.Now(),
				IssueInitialized: true,
				Issues:           issues,
			}

			filter := bson.M{"repository_id": repository.RepositoryID}
			update := bson.M{"$set": diff}
			_, err = repositoryCollection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				logrus.Warn(err)
			}
		}(repository)
	}
	wg.Wait()
	return nil
}

func fetchIssues(RepositoryName string) []*model.Issue {
	gitHubIssues, err := fetchGitHubIssues(RepositoryName)
	if err != nil {
		logrus.Error(err)
		return []*model.Issue{}
	}
	issues := make([]*model.Issue, 0, len(gitHubIssues))
	for _, gitHubIssue := range gitHubIssues {
		if gitHubIssue.HTMLURL != nil && strings.Contains(*gitHubIssue.HTMLURL, "pull") {
			continue
		}
		issues = append(issues, convertIssue(gitHubIssue))
	}
	return issues
}

func fetchGitHubIssues(repositoryName string) ([]*github.Issue, error) {
	ctx := context.Background()
	client := util.GetGitHubClient(ctx)
	repositoryOwner, repositoryName := strings.Split(repositoryName, "/")[0], strings.Split(repositoryName, "/")[1]
	listOpts := &github.ListOptions{Page: 1, PerPage: constant.ISSUES_API_RESULTS_PER_PAGE}
	opts := &github.IssueListByRepoOptions{State: "open", ListOptions: *listOpts}
	body, resp, _ := client.Issues.ListByRepo(ctx, repositoryOwner, repositoryName, opts)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d\n%v", resp.StatusCode, body)
	}
	return body, nil
}

func convertIssue(gitHubIssue *github.Issue) *model.Issue {
	issuer := new(model.User)
	if gitHubIssue.User != nil {
		issuer = convertUser(gitHubIssue.User)
	}

	labels := make([]*model.Label, 0)
	if gitHubIssue.Labels != nil {
		labels = convertLabels(gitHubIssue.Labels)
	}

	var assigneesCount *int
	if gitHubIssue.Assignees != nil {
		assigneesCountValue := len(gitHubIssue.Assignees)
		assigneesCount = &assigneesCountValue
	}

	issue := &model.Issue{
		IssueID:         *gitHubIssue.ID,
		GitHubCreatedAt: *gitHubIssue.CreatedAt,
		GitHubUpdatedAt: *gitHubIssue.UpdatedAt,
		Title:           *gitHubIssue.Title,
		URL:             *gitHubIssue.HTMLURL,
		AssigneesCount:  assigneesCount,
		CommentCount:    gitHubIssue.Comments,
		Issuer:          issuer,
		Labels:          labels,
	}
	return issue
}
