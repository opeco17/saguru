package update

import (
	"context"
	"fmt"
	"opeco17/saguru/job/constant"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/mongodb"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoNameAndID struct {
	RepositoryID int    `bson:"repository_id,omitempty"`
	Name         string `bson:"name,omitempty"`
}

func UpdateIssues(client *mongo.Client) error {
	logrus.Info("Start updating issues")

	repos, err := getRepositoriesFromMongoDB(client)
	if err != nil {
		logrus.Error("Failed to get repositories from MongoDB: %s", err.Error())
		return err
	}

	var wg sync.WaitGroup
	slots := make(chan interface{}, constant.UPDATE_ISSUE_SUBSET_SIZE)
	completed := make(chan interface{}, len(repos))

	wg.Add(len(repos))
	for _, repo := range repos {
		go func(repo *repoNameAndID) {
			slots <- struct{}{}
			defer wg.Done()
			defer func() { <-slots }()

			completed <- struct{}{}

			issues, err := fetchIssuesFromGitHub(repo.Name)
			if err != nil {
				logrus.Warn(fmt.Sprintf("Failed to fetch issues from GitHub: %s", err.Error()))
				return
			}
			if err := updateMongoDBIssues(issues, repo.RepositoryID, client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update issues in MongoDB: %s", err.Error()))
				return
			}
		}(repo)
	}

	go func() {
		counter := 0
		for range completed {
			if counter%100 == 0 {
				logrus.Info(fmt.Sprintf("Updating issues: %d/%d", counter, len(repos)))
			}
			counter += 1
		}
	}()

	wg.Wait()

	return nil
}

func getRepositoriesFromMongoDB(client *mongo.Client) ([]*repoNameAndID, error) {
	getSize := constant.UPDATE_ISSUE_SIZE
	notInitializedRepos, err := getRepositoriesSubsetFromMongoDB(client, getSize, false)
	if err != nil {
		logrus.Warn("Failed to get not initialized repositories: %s", err.Error())
		return nil, err
	}
	getSize = getSize - len(notInitializedRepos)

	initializedRepos := []*repoNameAndID{}
	if getSize > 0 {
		initializedRepos, err = getRepositoriesSubsetFromMongoDB(client, getSize, true)
		if err != nil {
			logrus.Warn("Failed to get initialized repositories: %s", err.Error())
			return nil, err
		}
	}

	repos := append(notInitializedRepos, initializedRepos...)
	return repos, nil
}

func getRepositoriesSubsetFromMongoDB(client *mongo.Client, maxGetSize int, isInitialized bool) ([]*repoNameAndID, error) {
	repoCollection := client.Database(mongodb.DATABASE_NAME).Collection(mongodb.REPOSITORY_COLLECTION_NAME)

	var repos []*repoNameAndID
	opts := options.Find().SetAllowDiskUse(true).SetLimit(int64(maxGetSize)).SetSort(bson.M{"updated_at": 1}).SetProjection(bson.M{"repository_id": 1, "name": 1})
	filter := bson.M{"issue_initialized": isInitialized}
	initializedIssueCursor, err := repoCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return []*repoNameAndID{}, err
	}
	if err = initializedIssueCursor.All(context.Background(), &repos); err != nil {
		return []*repoNameAndID{}, err
	}

	return repos, nil
}

func fetchIssuesFromGitHub(repoName string) ([]*mongodb.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.ISSUES_API_TIME_OUT)
	defer cancel()

	client := util.GetGitHubClient(ctx)

	repositoryOwner, repositoryName := strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1]
	listOpts := &github.ListOptions{Page: 1, PerPage: constant.ISSUES_API_RESULTS_PER_PAGE}
	opts := &github.IssueListByRepoOptions{State: "open", ListOptions: *listOpts}

	issues, resp, _ := client.Issues.ListByRepo(ctx, repositoryOwner, repositoryName, opts)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d: %v", resp.StatusCode, resp)
	}

	mongoDBIssues := make([]*mongodb.Issue, 0, len(issues))
	for _, issue := range issues {
		mongoDBIssues = append(mongoDBIssues, gitHubIssueToMongoDBIssue(issue))
	}

	return mongoDBIssues, nil
}

func updateMongoDBIssues(issues []*mongodb.Issue, repoID int, client *mongo.Client) error {
	repoCollection := client.Database(mongodb.DATABASE_NAME).Collection(mongodb.REPOSITORY_COLLECTION_NAME)

	diff := struct {
		UpdatedAt        time.Time        `bson:"updated_at"`
		IssueInitialized bool             `bson:"issue_initialized"`
		Issues           []*mongodb.Issue `bson:"issues"`
	}{
		UpdatedAt:        time.Now(),
		IssueInitialized: true,
		Issues:           issues,
	}

	filter := bson.M{"repository_id": repoID}
	update := bson.M{"$set": diff}
	if _, err := repoCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}

func gitHubIssueToMongoDBIssue(issue *github.Issue) *mongodb.Issue {
	issuer := new(mongodb.User)
	if issue.User != nil {
		issuer = gitHubUserToMongoDBUser(issue.User)
	}

	labels := make([]*mongodb.Label, 0)
	if issue.Labels != nil {
		labels = gitHubLabelsToMongoDBLabels(issue.Labels)
	}

	var assigneesCount *int
	if issue.Assignees != nil {
		assigneesCountValue := len(issue.Assignees)
		assigneesCount = &assigneesCountValue
	}

	converted := &mongodb.Issue{
		IssueID:         *issue.ID,
		GitHubCreatedAt: *issue.CreatedAt,
		GitHubUpdatedAt: *issue.UpdatedAt,
		Title:           *issue.Title,
		URL:             *issue.HTMLURL,
		AssigneesCount:  assigneesCount,
		CommentCount:    issue.Comments,
		Issuer:          issuer,
		Labels:          labels,
	}
	return converted
}

func gitHubUserToMongoDBUser(gitHubUser *github.User) *mongodb.User {
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

	user := &mongodb.User{
		UserID:    *gitHubUser.ID,
		Name:      name,
		URL:       url,
		AvatarURL: avatarURL,
	}
	return user
}

func gitHubLabelsToMongoDBLabels(gitHubLabels []*github.Label) []*mongodb.Label {
	labels := make([]*mongodb.Label, 0, len(gitHubLabels))
	for _, gitHubLabel := range gitHubLabels {
		name := ""
		if gitHubLabel.Name != nil {
			name = *gitHubLabel.Name
		}
		label := &mongodb.Label{
			LabelID: *gitHubLabel.ID,
			Name:    name,
		}
		labels = append(labels, label)
	}
	return labels
}
