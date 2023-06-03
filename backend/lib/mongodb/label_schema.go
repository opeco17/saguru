package mongodb

import "go.mongodb.org/mongo-driver/bson"

var labelSchema = bson.M{
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
}
