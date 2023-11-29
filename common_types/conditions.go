package types

type Condition struct {
	Value interface{} `validate:"" bson:"value,omitempty" json:"value,omitempty"`
	// it exists if there is a condition that checks equality
	Operator Operator `validate:"required,oneof=equals" json:"operator" bson:"operator"`
}

type KeyConditions struct {
	Conditions []Condition `json:"conditions" validate:"required,dive"`
	// With this set to true, only one of the conditions needs to be met
	OneOf bool `json:"oneOf"`
	// If set to true, will only run validation if it does happen to exist
	FieldOptional bool `json:"fieldOptional"`
}
