package types

type Operator string
type Location string

const (
	EQ Operator = "equals"
)

const (
	HEADER Location = "header"
)

type DocumentModel interface {
	GetCollectionName() string
}

type FilterableDocumentModel interface {
	DocumentModel
	FilterableProps() *[]string
}

type Action struct {
	Location Location
}

type CreatedIdResponse struct {
	ID string
}
