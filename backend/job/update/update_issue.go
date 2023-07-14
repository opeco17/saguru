package update

import (
	"context"
	"opeco17/saguru/job/constant"
	"opeco17/saguru/lib/mongodb"
	"sync"

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

	reposCh, err := getRepositoriesFromMongoDB(client)
	if err != nil {
		logrus.Error("Failed to get repositories from MongoDB: %s", err.Error())
		return err
	}

	for {
		var wg sync.WaitGroup

		concurrency := constant.UPDATE_ISSUE_SUBSET_SIZE
		if len(reposCh) < constant.UPDATE_ISSUE_SUBSET_SIZE {
			concurrency = len(reposCh)
		}
		wg.Add(concurrency)

		for i := 0; i < concurrency; i++ {
			go func() {
				defer wg.Done()
				repo := <-reposCh
				fetchIssuesFromGitHub(repo.Name)
			}()
		}
		wg.Wait()

		if len(reposCh) == 0 {
			break
		}
	}
	return nil
}

func getRepositoriesFromMongoDB(client *mongo.Client) (<-chan *repoNameAndID, error) {
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

	reposCh := make(chan *repoNameAndID, len(repos))
	for _, repo := range repos {
		reposCh <- repo
	}
	close(reposCh)
	return reposCh, nil
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

func fetchIssuesFromGitHub(repoName string) []*mongodb.Issue {
	return []*mongodb.Issue{}
}
