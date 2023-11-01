package conditions

import (
	"fmt"

	types "github.com/unexpectedtokens/api-tester/common"
)

func Validate(c *types.Condition, respBody map[string]string, key string) error {
	if c.Operator == types.EQ && respBody[key] != c.Value {
		return fmt.Errorf("expected %s to be equal to %s. Got %s instead", key, c.Value, respBody[key])
	}

	return nil
}

func ValidateConditions(r *types.TestCaseResult, respBody map[string]string) {

	for key, conditionsForKey := range r.Case.KeyConditions {

		for _, cond := range conditionsForKey.Conditions {
			err := Validate(&cond, respBody, key)
			if err != nil {
				r.AddErrMsg(fmt.Errorf("error running condition on key %s: %w", key, err).Error())
			}
		}

	}
}
