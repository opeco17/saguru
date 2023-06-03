package mongodb

type (
	User struct {
		UserID    int64  `bson:"user_id"`
		Name      string `bson:"name"`
		URL       string `bson:"url"`
		AvatarURL string `bson:"avatar_url"`
	}
)
