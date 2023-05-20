package schema

import "go.mongodb.org/mongo-driver/bson"

var issuerSchema = bson.M{
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
}
