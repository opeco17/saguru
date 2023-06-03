package mongodb

import "go.mongodb.org/mongo-driver/bson"

var RepositorySchema = bson.M{
	"bsonType": "object",
	"required": bson.A{
		"repository_id",
		"name",
		"url",
		"description",
		"star_count",
		"fork_count",
		"open_issue_count",
		"topics",
		"license",
		"language",
		"updated_at",
		"issue_initialized",
		"github_created_at",
		"github_updated_at",
		"issues",
	},
	"additionalProperties": true,
	"properties": bson.M{
		"repository_id": bson.M{
			"bsonType": bson.A{"int", "long"},
		},
		"name": bson.M{
			"bsonType": "string",
		},
		"url": bson.M{
			"bsonType": "string",
		},
		"description": bson.M{
			"bsonType": "string",
		},
		"star_count": bson.M{
			"bsonType": "int",
		},
		"fork_count": bson.M{
			"bsonType": "int",
		},
		"open_issue_count": bson.M{
			"bsonType": "int",
		},
		"topics": bson.M{
			"bsonType": "array",
			"items": bson.M{
				"bsonType": "string",
			},
		},
		"license": bson.M{
			"bsonType": "string",
		},
		"language": bson.M{
			"bsonType": "string",
		},
		"updated_at": bson.M{
			"bsonType": "date",
		},
		"issue_initialized": bson.M{
			"bsonType": "bool",
		},
		"github_created_at": bson.M{
			"bsonType": "date",
		},
		"github_updated_at": bson.M{
			"bsonType": "date",
		},
		"issues": bson.M{
			"bsonType": "array",
			"items":    issueSchema,
		},
	},
}
