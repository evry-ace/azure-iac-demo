package posts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorage struct {
	Coll *mongo.Collection
}

func (ps MongoStorage) GetPosts() []Post {
	ctx := context.TODO()

	var posts []Post
	cursor, err := ps.Coll.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &posts); err != nil {
		panic(err)
	}

	return posts
}

func (ps MongoStorage) SavePost(p Post) {
	ctx := context.TODO()
	_, err := ps.Coll.InsertOne(ctx, p)
	if err != nil {
		panic(err)
	}
}
