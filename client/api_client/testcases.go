package quikclient

import (
	"encoding/json"
	"fmt"
	"path"

	types "github.com/unexpectedtokens/api-tester/common_types"
)

func (q QuikClient) GetTestcases() ([]types.TestCase, error) {
	result, err := q.HTTPClient.Get(path.Join("/api/testcases", q.API_URL))

	if err != nil {
		return nil, fmt.Errorf("GetTescases performing get request: %w", err)
	}

	cases := []types.TestCase{}

	err = json.NewDecoder(result.Body).Decode(&cases)

	if err != nil {
		return nil, fmt.Errorf("GetTestcases decoding body: %w", err)
	}

	return cases, nil
}
