package data

import (
	"context"

	types "github.com/unexpectedtokens/api-tester/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateDocumentById[T types.DocumentModel](db *mongo.Database, ctx context.Context, objectID primitive.ObjectID, doc T) error {
	_, err := db.Collection(doc.GetCollectionName()).ReplaceOne(ctx, bson.M{"_id": objectID}, doc)

	return err
}
