package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestRequest struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Route string             `bson:"route" validate:"required" json:"route"`
	// Http verb "GET,POST..."
	Method string `bson:"method" json:"method" validate:"required,oneof=GET"`
}
