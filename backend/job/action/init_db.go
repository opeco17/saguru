package action

import (
	"context"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/mongodb"

	"github.com/sirupsen/logrus"
)

func InitMongoDB() error {
	logrus.Info("Start initializing DB.")

	client, err := util.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	mongodb.InitMongoDB(client)

	logrus.Info("Finished to initialize DB.")
	return nil
}
