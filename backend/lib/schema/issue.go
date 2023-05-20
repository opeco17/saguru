package schema

import "go.mongodb.org/mongo-driver/bson"

var issueSchema = bson.M{
	"bsonType": "object",
	"required": bson.A{
		"issue_id",
		"title",
		"url",
		"assignees_count",
		"comment_count",
		"github_created_at",
		"github_updated_at",
		"labels",
		"issuer",
	},
	"additionalProperties": true,
	"properties": bson.M{
		"issue_id": bson.M{
			"bsonType": bson.A{"int", "long"},
		},
		"title": bson.M{
			"bsonType": "string",
		},
		"url": bson.M{
			"bsonType": "string",
		},
		"assignees_count": bson.M{
			"bsonType": "int",
		},
		"comment_count": bson.M{
			"bsonType": "int",
		},
		"github_created_at": bson.M{
			"bsonType": "date",
		},
		"github_updated_at": bson.M{
			"bsonType": "date",
		},
		"labels": bson.M{
			"bsonType": "array",
			"items":    labelSchema,
		},
		"issuer": issuerSchema,
	},
}
