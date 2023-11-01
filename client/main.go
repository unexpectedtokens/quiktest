package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/unexpectedtokens/api-tester/client/report"
	"github.com/unexpectedtokens/api-tester/client/test_request"
	types "github.com/unexpectedtokens/api-tester/common"
)

const DEFAULT_URL = "http://localhost:8080/testcases"

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
	err := test_request.Ping(DEFAULT_URL)
	if err != nil {
		panic(err)
	}
	resp, err := http.Get(DEFAULT_URL)

	if err != nil {
		panic(err)
	}

	result := types.TestReport{
		Title:             fmt.Sprintf("Testresult from %s", time.Now().Format(time.Layout)),
		TotalTestDuration: time.Hour * 3,
	}

	cases := []types.TestCase{}

	err = json.NewDecoder(resp.Body).Decode(&cases)

	if err != nil {
		panic(err)
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
