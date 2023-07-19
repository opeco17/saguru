package mongodb

import (
	"context"
	"fmt"
	"strconv"

	errorsutil "opeco17/saguru/lib/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient(user string, password string, host string, port int) (*mongo.Client, error) {
	if host == "" {
		return nil, errorsutil.CustomError{Message: "You must set 'host' to connect to MongoDB"}
	}
	if port < 0 {
		return nil, errorsutil.CustomError{Message: "Port number is invalid"}
	}
	if user == "" {
		return nil, errorsutil.CustomError{Message: "You must set 'user' to connect to MongoDB"}
	}
	if password == "" {
		return nil, errorsutil.CustomError{Message: "You must set 'password' to connect to MongoDB"}
	}

	credential := options.Credential{
		Username: user,
		Password: password,
	}
	clientOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s", host, strconv.Itoa(port))).SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}
	return client, nil
}
