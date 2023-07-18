package main

import (
	"context"
	"fmt"
	"opeco17/saguru/job/update"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/mongodb"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update resources in database.",
	Run: func(cmd *cobra.Command, args []string) {
		options, err := update.GetOptions(cmd.Flags())
		if err != nil {
			logrus.Error("Failed to parse flags")
			os.Exit(1)
		}

		if !options.Issue && !options.Repository && !options.Cache && !options.Index {
			logrus.Error("One of the target to update should be specified")
			os.Exit(1)
		}

		mongoDBClient, err := util.GetMongoDBClient()
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to connect to MongoDB: %s", err.Error()))
			os.Exit(1)
		}
		defer mongoDBClient.Disconnect(context.Background())
		mongodb.InitMongoDB(mongoDBClient)
		logrus.Info("Finished to initialize database")

		memcachedClient, err := util.GetMemcachedClient()
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to connect to Memcached: %s", err.Error()))
			os.Exit(1)
		}
		defer memcachedClient.Close()

		if options.Repository {
			if err := update.UpdateRepositories(mongoDBClient); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update repositories: %s", err.Error()))
			}
		}
		if options.Issue {
			if err := update.UpdateIssues(mongoDBClient); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update issues: %s", err.Error()))
			}
		}
		if options.Cache {
			if err := update.UpdateCaches(mongoDBClient, memcachedClient); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update caches: %s", err.Error()))
			}
		}
		if options.Cache {
			if err := update.UpdateCaches(mongoDBClient, memcachedClient); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update caches: %s", err.Error()))
			}
		}
		if options.Index {
			if err := update.UpdateIndices(mongoDBClient); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update indices: %s", err.Error()))
			}
		}
	},
}

func init() {
	updateCmd.Flags().BoolP("issue", "", false, "Issues are updated when specified")
	updateCmd.Flags().BoolP("repository", "", false, "Repositories are updated when specified")
	updateCmd.Flags().BoolP("cache", "", false, "caches are updated when specified")
	updateCmd.Flags().BoolP("index", "", false, "indices are updated when specified")
	rootCmd.AddCommand(updateCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
