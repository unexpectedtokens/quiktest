package testrunner

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	quikclient "github.com/unexpectedtokens/api-tester/client/api_client"
	types "github.com/unexpectedtokens/api-tester/common_types"
)

type ApiClient interface {
}

type TestRunner struct {
	QuikClient ApiClient
	TestClient http.Client
}

func testSingleCase(tCase *types.TestCase, ch <-chan *types.TestCaseResult) {

}

func (t *TestRunner) testGroup(in <-chan types.TestCase) <-chan *types.TestCaseResult {
	out := make(chan *types.TestCaseResult)

	// Groups should run concurrently to drastically reduce execution time.
	// However, don't run the tests itself concurrently, there could be a desired order of execution.
	// Make this a setting?
	go func() {
		for tCase := range in {
			fmt.Println("RUNNING", tCase.ID.Hex(), tCase.Title)
			// Run test and send output over channel
			result := t.SendTestRequest(tCase)

			fmt.Println(result.Case.ID.Hex(), result.ErrMessages)
			out <- result
		}

		close(out)
	}()

	return out
}

func mergeChannels(chans ...<-chan *types.TestCaseResult) <-chan *types.TestCaseResult {

	var wg sync.WaitGroup
	out := make(chan *types.TestCaseResult)
	output := func(c <-chan *types.TestCaseResult) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chans))

	for _, c := range chans {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(tCases []types.TestCase) <-chan types.TestCase {
	out := make(chan types.TestCase)
	go func() {
		for _, tCase := range tCases {
			out <- tCase
		}
		close(out)
	}()
	return out
}

func (t *TestRunner) RunTestPipeline(groupedCases *quikclient.GroupedCases) <-chan *types.TestCaseResult {

	chans := []<-chan *types.TestCaseResult{}
	for groupID, testcases := range *groupedCases {
		log.Printf("Starting testrun for group %s. It has %d cases\n", groupID, len(testcases))
		// TODO: make amount of goroutines configurable
		chans = append(chans, t.testGroup(gen(testcases)))
	}

	out := mergeChannels(chans...)

	return out
}
