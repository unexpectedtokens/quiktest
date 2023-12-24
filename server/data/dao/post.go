package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *DAO[T]) SaveDocument(doc T, ctx context.Context) (insertedId string, err error) {
	collectionName := doc.GetCollectionName()
	result, err := d.db.Collection(collectionName).InsertOne(ctx, doc)

	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)

	if ok {
		return id.Hex(), nil
	}

	return "", nil
}
