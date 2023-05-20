package model

type (
	Label struct {
		LabelID int64  `bson:"label_id"`
		Name    string `bson:"name"`
	}
)
