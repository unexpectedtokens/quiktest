package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dejong543:kQimHN27!V_gBep@tester.cyn9pbc.mongodb.net/?retryWrites=true&w=majority"))
	return client, err
}

func GetDB(client *mongo.Client) *mongo.Database {
	return client.Database("testcases")
}
