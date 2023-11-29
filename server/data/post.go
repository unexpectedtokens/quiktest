package data

import (
	"context"

	"github.com/rs/zerolog"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveDocument[T types.DocumentModel](db *mongo.Database, doc T, ctx context.Context, logger zerolog.Logger) (insertedId string, err error) {
	collectionName := doc.GetCollectionName()
	result, err := db.Collection(collectionName).InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)

	if ok {
		return id.Hex(), nil
	}

	return "", nil
}
