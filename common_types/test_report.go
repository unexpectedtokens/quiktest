package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestReport struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title                string             `json:"title" bson:"title" validate:"required"`
	TestCaseResults      []TestCaseResult   `json:"-" bson:"-"`
	SuccessPercentage    float32            `json:"successPercentage" bson:"successPercentage" validate:"min=0,max=100"`
	TotalTestDuration    time.Duration      `json:"totalTestDuration" bson:"totalTestDuration"`
	TotalTestDurationFmt string             `bson:"-" json:"totalTestDurationFmt,omitempty"`
}

func (T TestReport) GetCollectionName() string {
	return "test_reports"
}

var filterableTestReportProps *[]string = &[]string{}

type TestCaseResult struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Case             TestCase           `bson:"case" json:"case"`
	ActualReturnCode int                `bson:"actualReturnCode" json:"actualReturnCode"`
	ResponseTime     time.Duration      `bson:"responseTime" json:"responseTime"`
	ResponseBody     string             `bson:"responseBody" json:"responseBody"`
	TestReportId     primitive.ObjectID `bson:"testReportId" json:"testReportId" validate:"required"`
	ErrMessages      []string           `bson:"errorMessages" json:"errorMessages"`
}

var filterableTestCaseResultProps *[]string = &[]string{
	"testReportId",
}

func FilterableTestCaseResultProps() *[]string {
	return filterableTestCaseResultProps
}

func (T *TestCaseResult) AddErrMsg(msg string) {
	T.ErrMessages = append(T.ErrMessages, msg)
}

func (T TestCaseResult) GetCollectionName() string {
	return "test_case_results"
}
