package data

import (
	"context"
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common_types"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteDocument[T types.DocumentModel](db *mongo.Database, collection string, ctx context.Context, query interface{}) error {
	dbCollection := db.Collection(collection)

	deletedCount, err := dbCollection.DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	if deletedCount.DeletedCount == 0 {
		return fmt.Errorf("error deleting document: no match found for query %a", query)
	}

	return nil
}
