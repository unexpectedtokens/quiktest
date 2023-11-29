package quikclient

import (
	"encoding/json"
	"fmt"
	"path"

	types "github.com/unexpectedtokens/api-tester/common_types"
)

func (q QuikClient) Groups() ([]types.TestGroup, error) {

	resp, err := q.HTTPClient.Get(path.Join(q.API_URL, "/api/testgroups"))

	if err != nil {
		return nil, fmt.Errorf("error getting groups: %w", err)
	}

	groups := []types.TestGroup{}

	err = json.NewDecoder(resp.Body).Decode(&groups)

	return groups, err
}
