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

		if !options.Issue && !options.Repository && !options.Cache {
			logrus.Error("One of the target to update should be specified")
			os.Exit(1)
		}

		client, err := util.GetMongoDBClient()
		if err != nil {
			logrus.Error("Failed to connect to database")
			os.Exit(1)
		}
		defer client.Disconnect(context.Background())

		mongodb.InitMongoDB(client)
		logrus.Info("Finished to initialize database")

		if options.Repository {
			if err := update.UpdateRepositories(client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update repositories: %s", err.Error()))
			}
		}
		if options.Issue {
			if err := update.UpdateIssues(client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update issues: %s", err.Error()))
			}
		}
		if options.Cache {
			if err := update.UpdateCache(client); err != nil {
				logrus.Warn(fmt.Sprintf("Failed to update caches: %s", err.Error()))
			}
		}
	},
}

func init() {
	updateCmd.Flags().BoolP("issue", "", false, "Issues are updated when specified")
	updateCmd.Flags().BoolP("repository", "", false, "Repositories are updated when specified")
	updateCmd.Flags().BoolP("cache", "", false, "caches are updated when specified")
	rootCmd.AddCommand(updateCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
