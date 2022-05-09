package lib

import "go.mongodb.org/mongo-driver/bson"

var MongoSchema = bson.M{
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
			"items": bson.M{
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
						"items": bson.M{
							"bsonType": "object",
							"required": bson.A{
								"label_id",
								"name",
							},
							"additionalProperties": true,
							"properties": bson.M{
								"label_id": bson.M{
									"bsonType": bson.A{"int", "long"},
								},
								"name": bson.M{
									"bsonType": "string",
								},
							},
						},
					},
					"issuer": bson.M{
						"bsonType": "object",
						"required": bson.A{
							"user_id",
							"name",
							"url",
							"avatar_url",
						},
						"additionalProperties": true,
						"properties": bson.M{
							"user_id": bson.M{
								"bsonType": bson.A{"int", "long"},
							},
							"name": bson.M{
								"bsonType": "string",
							},
							"url": bson.M{
								"bsonType": "string",
							},
							"avatar_url": bson.M{
								"bsonType": "string",
							},
						},
					},
				},
			},
		},
	},
}
