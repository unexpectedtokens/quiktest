package quikclient

import (
	"encoding/json"
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common_types"
)

func (q QuikClient) groups() ([]types.TestGroup, error) {

	resp, err := q.HTTPClient.Get(q.formatUrl("/testgroups"))

	if err != nil {
		return nil, fmt.Errorf("error getting groups: %w", err)
	}

	groups := []types.TestGroup{}

	err = json.NewDecoder(resp.Body).Decode(&groups)

	return groups, err
}
