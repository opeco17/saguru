package model

type (
	CachedItem struct {
		Name  string `bson:"name"`
		Count int    `bson:"count"`
	}
)
