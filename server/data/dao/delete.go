package dao

import (
	"context"
	"fmt"
)

func (d *DAO[T]) DeleteDocument(collection string, ctx context.Context, query interface{}) error {
	dbCollection := d.db.Collection(collection)

	deletedCount, err := dbCollection.DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	if deletedCount.DeletedCount == 0 {
		return fmt.Errorf("error deleting document: no match found for query %a", query)
	}

	return nil
}
