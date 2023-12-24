package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *DAO[T]) UpdateDocumentById(ctx context.Context, objectID primitive.ObjectID, doc T) error {
	_, err := d.db.Collection(doc.GetCollectionName()).ReplaceOne(ctx, bson.M{"_id": objectID}, doc)

	return err
}
