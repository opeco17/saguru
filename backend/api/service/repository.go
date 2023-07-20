package service

import (
	"context"
	"fmt"
	"opeco17/saguru/api/constant"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	errorsutil "opeco17/saguru/lib/errors"
	"opeco17/saguru/lib/mongodb"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRepositoriesFromMongoDB(client *mongo.Client, input *model.GetRepositoriesInput) ([]mongodb.Repository, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	repositoryCollection := client.Database("main").Collection("repositories")
	filter := bson.M{}

	// Filter about repositories
	if input.Languages != "" {
		filter["language"] = bson.M{"$in": strings.Split(input.Languages, ",")}
	}
	if input.License != "" {
		filter["license"] = input.License
	}
	if input.Keyword != "" {
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{"$regex": input.Keyword, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": input.Keyword, "$options": "i"}},
		}
	}
	if input.StarCountLower != nil || input.StarCountUpper != nil {
		starCountFilter := bson.M{}
		if input.StarCountLower != nil {
			starCountFilter["$gte"] = input.StarCountLower
		}
		if input.StarCountUpper != nil {
			starCountFilter["$lte"] = input.StarCountUpper
		}
		filter["star_count"] = starCountFilter
	}
	if input.ForkCountLower != nil || input.ForkCountUpper != nil {
		forkCountFilter := bson.M{}
		if input.ForkCountLower != nil {
			forkCountFilter["$gte"] = input.ForkCountLower
		}
		if input.StarCountUpper != nil {
			forkCountFilter["$lte"] = input.ForkCountUpper
		}
		filter["fork_count"] = forkCountFilter
	}

	// Filter about issues
	issueFilter := bson.M{"assignees_count": bson.M{"$gte": 0}} // To remove empty issues
	if input.IsAssigned != nil && *input.IsAssigned {
		issueFilter["assignees_count"] = bson.M{"$gte": 1}
	} else if input.IsAssigned != nil && !*input.IsAssigned {
		issueFilter["assignees_count"] = 0
	}
	if input.Labels != "" {
		issueFilter["labels.name"] = bson.M{"$in": strings.Split(input.Labels, ",")}
	}
	filter["issues"] = bson.M{"$elemMatch": issueFilter, "$exists": true}

	logrus.Info(fmt.Sprintf("Filter %+v", filter))

	// Set options
	var metric string
	var direction int
	orderBy := input.Orderby
	if orderBy == "" {
		orderBy = "STAR_COUNT_DESC"
	}
	if strings.Contains(orderBy, "STAR_COUNT") {
		metric = "star_count"
	} else if strings.Contains(orderBy, "FORK_COUNT") {
		metric = "fork_count"
	}
	if strings.Contains(orderBy, "DESC") {
		direction = -1
	} else if strings.Contains(orderBy, "ASC") {
		direction = 1
	}
	opts := options.Find().SetLimit(int64(constant.RESULTS_PER_PAGE + 1)).SetSkip(int64(constant.RESULTS_PER_PAGE * input.Page)).SetSort(bson.M{metric: direction})

	cursor, err := repositoryCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, errorsutil.Wrap(err, "Failed to get repositories from MongoDB")
	}
	var repositories []mongodb.Repository
	if err = cursor.All(context.Background(), &repositories); err != nil {
		return nil, errorsutil.Wrap(err, "Failed to get repositories from MongoDB")
	}
	return repositories, nil
}

func FilterIssuesInRepositories(repositories []mongodb.Repository, input *model.GetRepositoriesInput) []mongodb.Repository {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	filteredRepositories := make([]mongodb.Repository, 0, len(repositories))
	assigneeFilter := func(assigneesCount int) bool { return true }
	labelFilter := func(labels []*mongodb.Label) bool { return true }

	// Set filter
	if input.IsAssigned != nil && *input.IsAssigned {
		assigneeFilter = func(assigneesCount int) bool { return assigneesCount > 0 }
	} else if input.IsAssigned != nil && !*input.IsAssigned {
		assigneeFilter = func(assigneesCount int) bool { return assigneesCount == 0 }
	}

	if input.Labels != "" {
		labelFilter = func(labels []*mongodb.Label) bool {
			inputLabelNames := strings.Split(input.Labels, ",")
			for _, label := range labels {
				for _, inputLabelName := range inputLabelNames {
					if label.Name == inputLabelName {
						return true
					}
				}
			}
			return false
		}
	}

	// Filter issues
	for _, repository := range repositories {
		filteredIssues := make([]*mongodb.Issue, 0, len(repository.Issues))
		for _, issue := range repository.Issues {
			if assigneeFilter(*issue.AssigneesCount) && labelFilter(issue.Labels) {
				filteredIssues = append(filteredIssues, issue)
			}
		}
		repository.Issues = filteredIssues
		filteredRepositories = append(filteredRepositories, repository)
	}
	return filteredRepositories
}
