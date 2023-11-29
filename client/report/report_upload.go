package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	types "github.com/unexpectedtokens/api-tester/common_types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const API_URL = "http://localhost:8080/api"

func CreateTestReport(report types.TestReport) (string, error) {
	jsonOutput, err := json.Marshal(report)

	if err != nil {
		panic(err)
	}
	resp, err := http.Post(fmt.Sprintf("%s/testreports", API_URL), "application/json", strings.NewReader(string(jsonOutput)))

	if err != nil {
		log.Printf("error sending post request, status %d: %s", resp.StatusCode, err.Error())
		return "", err
	} else if resp.StatusCode == http.StatusUnprocessableEntity {

		return "", errors.New("a validation error occurred")
	}

	respBody := types.CreatedIdResponse{}

	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		log.Printf("error decoding json response body: %s", err.Error())
		return "", err
	}

	return respBody.ID, nil
}

func CreateTestCaseResults(results []types.TestCaseResult, reportID string) {
	objectId, err := primitive.ObjectIDFromHex(reportID)

	if err != nil {
		panic(err)
	}
	totalCaseResults := len(results)
	for i, testCaseResult := range results {
		log.Printf("uploading testresult %d/%d...", i+1, totalCaseResults)

		testCaseResult.TestReportId = objectId
		jsonPayload, err := json.Marshal(testCaseResult)

		if err != nil {
			log.Printf("error generating json payload: %s", err.Error())
			continue
		}

		resp, err := http.Post(fmt.Sprintf("%s/testreports/results", API_URL), "application/json", strings.NewReader(string(jsonPayload)))

		if err != nil {
			log.Printf("error uploading testcase result: %s", err.Error())
			continue
		}

		log.Printf("Uploading of testresult %d/%d complete, statuscode: %s", i+1, totalCaseResults, resp.Status)
	}

}
