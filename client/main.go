package client

import (
	"fmt"
	"log"
	"net/http"
	"time"

	quikclient "github.com/unexpectedtokens/api-tester/client/api_client"
	"github.com/unexpectedtokens/api-tester/client/report"
	"github.com/unexpectedtokens/api-tester/client/test_request"
	types "github.com/unexpectedtokens/api-tester/common_types"
)

const DEFAULT_URL = "http://localhost:8080"

// TODO: Create a pipeline out of this

func RunTests(cases []types.TestCase, url string) (testCaseResults []types.TestCaseResult, totalTestDuration time.Duration, succesPercentage float32) {
	clnt := http.Client{}

	testCaseResults = []types.TestCaseResult{}

	totalTestStartTime := time.Now()

	totalCases := float32(len(cases))
	var passingCases int
	for i, testCase := range cases {
		log.Printf("running test %d/%d\n", i+1, int(totalCases))

		testCaseResult := test_request.SendTestRequest(&clnt, testCase)

		fmt.Println(testCaseResult.ErrMessages)
		if len(testCaseResult.ErrMessages) == 0 {
			passingCases += 1
			fmt.Println("Great success")
		}
		testCaseResults = append(testCaseResults, testCaseResult)
	}

	succesPercentage = float32(passingCases) / totalCases * 100
	totalTestDuration = time.Since(totalTestStartTime)

	return
}

func RunClient() {

	quikClient := quikclient.QuikClient{
		API_URL:    DEFAULT_URL,
		HTTPClient: http.Client{},
	}

	err := quikClient.Ping()
	if err != nil {
		panic(err)
	}

	cases, err := quikClient.GetTestcases()
	if err != nil {
		panic(err)
	}

	result := types.TestReport{
		Title:             fmt.Sprintf("Testresult from %s", time.Now().Format(time.Layout)),
		TotalTestDuration: time.Hour * 3,
	}

	var testCaseResults []types.TestCaseResult

	testCaseResults, totalTestRunDuration, succesPercentage := RunTests(cases, "")
	result.SuccessPercentage = succesPercentage
	result.TotalTestDuration = totalTestRunDuration
	reportID, err := report.CreateTestReport(result)

	if err != nil {
		panic(err)
	}

	report.CreateTestCaseResults(testCaseResults, reportID)

	log.Printf("finished creating testreport. Testreport id: %s. Title: %s, successPercentage %.2f", reportID, result.Title, result.SuccessPercentage)
}
