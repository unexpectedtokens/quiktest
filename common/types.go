package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operator string
type Location string

const (
	EQ Operator = "equals"
)

const (
	HEADER Location = "header"
)

type DocumentModel interface {
	GetCollectionName() string
}

type Action struct {
	Location Location
}
type Condition struct {
	Value interface{} `validate:"" bson:"value,omitempty" json:"value,omitempty"`
	// it exists if there is a condition that checks equality
	Operator Operator `validate:"required,oneof=equals" json:"operator" bson:"operator"`
}

type KeyConditions struct {
	Conditions []Condition `json:"conditions" validate:"required,dive"`
	// With this set to true, only one of the conditions needs to be met
	OneOf bool `json:"oneOf"`
	// If set to true, will only run validation if it does happen to exist
	FieldOptional bool `json:"fieldOptional"`
}

type TestRequest struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Route string             `bson:"route" validate:"required" json:"route"`
	// Http verb "GET,POST..."
	Method string `bson:"method" json:"method" validate:"required,oneof=GET POST"`
}

type TestCase struct {
	ID               primitive.ObjectID       `bson:"_id,omitempty" json:"_id,omitempty"`
	Title            string                   `bson:"title" json:"title" validate:"required"`
	Request          TestRequest              `bson:"request" json:"request"`
	RequestID        primitive.ObjectID       `bson:"request_id" json:"request_id"`
	ExpectReturnCode int                      `bson:"expected_return_code" validate:"required" json:"expectReturnCode"`
	KeyConditions    map[string]KeyConditions `bson:"keyConditions" json:"keyConditions" validate:"dive"`
}

type TestCaseResult struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Case             TestCase           `bson:"case" json:"case"`
	ActualReturnCode int                `bson:"actualReturnCode" json:"actualReturnCode"`
	ResponseTime     time.Duration      `bson:"responseTime" json:"responseTime"`
	ResponseBody     string             `bson:"responseBody" json:"responseBody"`
	TestReportId     primitive.ObjectID `bson:"testReportId" json:"testReportId"`
	ErrMessages      []string           `bson:"errorMessages" json:"errorMessages"`
}

type TestReport struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title                string             `json:"title" bson:"title" validate:"required"`
	TestCaseResults      []TestCaseResult   `json:"-" bson:"-"`
	SuccessPercentage    float32            `json:"successPercentage" bson:"successPercentage" validate:"min=0,max=100"`
	TotalTestDuration    time.Duration      `json:"totalTestDuration" bson:"totalTestDuration"`
	TotalTestDurationFmt string             `bson:"-" json:"totalTestDurationFmt,omitempty"`
}

type CreatedIdResponse struct {
	ID string
}

func (T *TestCaseResult) AddErrMsg(msg string) {
	T.ErrMessages = append(T.ErrMessages, msg)
}

func (T TestCase) GetCollectionName() string {
	return "test_cases"
}

func (T TestReport) GetCollectionName() string {
	return "test_reports"
}

func (T TestCaseResult) GetCollectionName() string {
	return "test_case_results"
}
