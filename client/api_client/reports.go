package quikclient

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

func (q QuikClient) CreateTestReport(report types.TestReport) (string, error) {
	jsonOutput, err := json.Marshal(report)

	if err != nil {
		panic(err)
	}

	url := q.formatUrl("/testreports")
	resp, err := q.HTTPClient.Post(url, "application/json", strings.NewReader(string(jsonOutput)))

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

func (q QuikClient) CreateTestCaseResultsFromChannel(resultsChannel <-chan *types.TestCaseResult, reportID string) {
	objectId, err := primitive.ObjectIDFromHex(reportID)

	if err != nil {
		panic(err)
	}
	for testCaseResult := range resultsChannel {
		log.Printf("uploading testresult for case %s... ", testCaseResult.Case.ID.Hex())

		testCaseResult.TestReportId = objectId
		jsonPayload, err := json.Marshal(testCaseResult)

		if err != nil {
			log.Printf("[ERROR]: generating json payload: %s\n", err.Error())
			continue
		}

		_, err = q.HTTPClient.Post(q.formatUrl("/testreports/results"), "application/json", strings.NewReader(string(jsonPayload)))

		if err != nil {
			log.Printf("[ERROR]: %s\n", err.Error())
			continue
		}

		fmt.Print("[DONE]\n")
	}
}
