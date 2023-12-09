package quikclient

import (
	"encoding/json"
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common_types"
)

const UNGROUPED = "ungrouped"

type GroupedCases map[string][]types.TestCase

func (q QuikClient) GetTestcases() (*[]types.TestCase, error) {
	url := q.formatUrl("/testcases")
	result, err := q.HTTPClient.Get(url)

	if err != nil {
		return nil, fmt.Errorf("GetTescases performing get request: %w", err)
	}

	cases := []types.TestCase{}

	err = json.NewDecoder(result.Body).Decode(&cases)

	if err != nil {
		return nil, fmt.Errorf("GetTestcases decoding body: %w", err)
	}

	return &cases, nil
}

func (q QuikClient) SortByGroup(cases *[]types.TestCase) (*GroupedCases, error) {
	gc := GroupedCases{
		UNGROUPED: []types.TestCase{},
	}

	groups, err := q.groups()

	if err != nil {
		return nil, err
	}

	for _, tCase := range *cases {
		if tCase.TestGroupID == nil {
			gc[UNGROUPED] = append(gc[UNGROUPED], tCase)
		} else {
			groupFound := false
			for _, group := range groups {
				groupID, tCaseGroupID := group.ID.Hex(), tCase.TestGroupID.Hex()

				if groupID == (tCaseGroupID) {
					if _, ok := gc[groupID]; !ok {
						gc[groupID] = []types.TestCase{}
					}
					groupFound = true
					gc[groupID] = append(gc[groupID], tCase)
				}

			}

			if !groupFound {
				gc[UNGROUPED] = append(gc[UNGROUPED], tCase)
			}
		}
	}

	return &gc, nil
}
