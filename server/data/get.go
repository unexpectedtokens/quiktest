package data

import (
	"context"
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDocument[T types.DocumentModel](db *mongo.Database, collection string, query interface{}, ctx context.Context) (T, error) {
	var document T
	err := db.Collection(collection).FindOne(ctx, query).Decode(&document)

	if err != nil {
		return document, fmt.Errorf("error finding one document: %w", err)
	}

	return document, nil
}

func GetDocuments[T types.DocumentModel](db *mongo.Database, collection string, query interface{}, ctx context.Context) ([]T, error) {
	cursor, err := db.Collection(collection).Find(ctx, query)

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
