package mongodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient(user string, password string, host string, port int) (*mongo.Client, error) {
	if host == "" {
		err_message := "You must set 'host' to connect to MongoDB"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}
	if port < 0 {
		err_message := "Port number is invalid"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}
	if user == "" {
		err_message := "You must set 'user' to connect to MongoDB"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}
	if password == "" {
		err_message := "You must set 'password' to connect to MongoDB"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}

	credential := options.Credential{
		Username: user,
		Password: password,
	}
	clientOpts := options.Client().ApplyURI("mongodb://" + host + ":" + strconv.Itoa(port)).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}
