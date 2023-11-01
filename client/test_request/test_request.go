package test_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/unexpectedtokens/api-tester/client/conditions"
	types "github.com/unexpectedtokens/api-tester/common"
)

func SendTestRequest(client *http.Client, testCase types.TestCase) types.TestCaseResult {
	now := time.Now()

	testCaseResult := types.TestCaseResult{
		Case: testCase,
	}

	url := fmt.Sprintf("http://localhost:8181%s", testCase.Request.Route)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		testCaseResult.AddErrMsg(fmt.Sprintf("Error constructing request: %s", err.Error()))
		return testCaseResult
	}

	resp, err := client.Do(req)
	if err != nil {
		testCaseResult.AddErrMsg(fmt.Sprintf("error sending request: %s", err.Error()))
		return testCaseResult
	}

	timeToRespond := time.Since(now)

	testCaseResult.ResponseTime = timeToRespond

	testCaseResult.ActualReturnCode = resp.StatusCode
	if testCase.ExpectReturnCode != resp.StatusCode {
		testCaseResult.AddErrMsg(fmt.Sprintf("Expected statuscode %d, got %d instead", testCase.ExpectReturnCode, resp.StatusCode))
	}

	respBody := make(map[string]string)

	if testCase.KeyConditions != nil {
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			testCaseResult.AddErrMsg("Expected body to run conditions on, but found EOF")
		} else {
			conditions.ValidateConditions(&testCaseResult, respBody)
		}
	}

	if len(respBody) > 0 {
		respBodyAsString, err := json.Marshal(respBody)
		if err == nil {
			testCaseResult.ResponseBody = string(respBodyAsString)
		}
	}

	return testCaseResult
}
