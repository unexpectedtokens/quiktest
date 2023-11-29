package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestGroup struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string             `bson:"title" json:"title" validate:"required"`
}

func (T TestGroup) GetCollectionName() string {
	return "test_groups"
}

var filterableTestGroupProps *[]string = &[]string{}

func (T TestGroup) FilterableProps() *[]string {
	return filterableTestGroupProps
}
