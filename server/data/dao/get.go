package dao

import (
	"context"
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common_types"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *DAO[T]) GetDocument(collection string, query interface{}, ctx context.Context) (types.DocumentModel, error) {
	var document T
	err := d.db.Collection(collection).FindOne(ctx, query).Decode(&document)

	if err != nil {
		return document, fmt.Errorf("error finding one document: %w", err)
	}

	return document, nil
}

func (d *DAO[T]) GetDocuments(collection string, query interface{}, ctx context.Context) ([]T, error) {
	cursor, err := d.db.Collection(collection).Find(ctx, query)

	documents := []T{}
	if err != nil {
		if err == mongo.ErrNilDocument {
			return documents, nil
		}
		return nil, fmt.Errorf("error querying all testcases: %w", err)
	}

	err = cursor.All(ctx, &documents)

	if err != nil {
		return nil, fmt.Errorf("error converting cursor to documents: %w", err)
	}

	return documents, nil
}
