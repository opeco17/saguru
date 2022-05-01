package main

import (
	"fmt"
	"opeco17/gitnavi/lib"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func getMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")

	client, err := lib.GetMongoDBClient(user, password, host)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}

func orderMetrics() []string {
	return []string{
		"star_count",
		"fork_count",
	}
}

func setOrderQuery(query *gorm.DB, orderby string) {
	order_query := "repositories.id"
	for _, metric := range orderMetrics() {
		if orderby == metric+"_asc" {
			order_query = fmt.Sprintf("repositories.%s ASC", metric)
			break
		} else if orderby == metric+"_desc" {
			order_query = fmt.Sprintf("repositories.%s DESC", metric)
			break
		}
	}
	query.Order(order_query)
}

func setDistinctQuery(query *gorm.DB, orderby string) {
	distinct_query := "repositories.id"
	for _, metric := range orderMetrics() {
		if orderby == metric+"_asc" {
			distinct_query = fmt.Sprintf("repositories.id, repositories.%s", metric)
			break
		} else if orderby == metric+"_desc" {
			distinct_query = fmt.Sprintf("repositories.id, repositories.%s", metric)
			break
		}
	}
	query.Distinct(distinct_query)
}
