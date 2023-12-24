package dao

import (
	types "github.com/unexpectedtokens/api-tester/common_types"
	"go.mongodb.org/mongo-driver/mongo"
)

type DAO[T types.DocumentModel] struct {
	db *mongo.Database
}

func New[T types.DocumentModel](db *mongo.Database) *DAO[T] {
	return &DAO[T]{
		db: db,
	}
}
