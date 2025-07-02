package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func generateID(ctx context.Context, collection *mongo.Collection) (int64, error) {
	filter := bson.M{"_id": "global_id"}
	update := bson.M{"$inc": bson.M{"curr": 1}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.Before)

	var result struct {
		Curr int64 `bson:"curr"`
	}

	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	return result.Curr + 1, nil
}
