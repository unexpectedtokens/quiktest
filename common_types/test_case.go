package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestCase struct {
	ID               primitive.ObjectID       `bson:"_id,omitempty" json:"_id,omitempty"`
	Title            string                   `bson:"title" json:"title" validate:"required"`
	Request          TestRequest              `bson:"request" json:"request"`
	RequestID        primitive.ObjectID       `bson:"request_id" json:"request_id"`
	ExpectReturnCode int                      `bson:"expected_return_code" validate:"required" json:"expectReturnCode"`
	KeyConditions    map[string]KeyConditions `bson:"keyConditions" json:"keyConditions" validate:"dive"`
	TestGroupID      *primitive.ObjectID      `bson:"testGroup,omitempty" json:"testGroup,omitempty"`
}

func (T TestCase) GetCollectionName() string {
	return "test_cases"
}

var filterableTestcaseProps *[]string = &[]string{
	"testGroup",
}

func (T TestCase) FilterableProps() *[]string {
	return filterableTestcaseProps
}
