package lib

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient(user string, password string, host string) (*mongo.Client, error) {
	if host == "" {
		message := "You must set 'host' to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
	}
	if user == "" {
		message := "You must set 'user' to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
	}
	if password == "" {
		message := "You must set 'password' to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
	}

	credential := options.Credential{
		Username: user,
		Password: password,
	}
	clientOpts := options.Client().ApplyURI("mongodb://" + host).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}
