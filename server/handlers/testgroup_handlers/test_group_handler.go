package testgrouphandlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"github.com/unexpectedtokens/api-tester/server/handlers/generic_handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestGroupHandler struct {
	db       *mongo.Database
	router   chi.Router
	validate *validator.Validate
}

func (t *TestGroupHandler) Initialize(db *mongo.Database, router chi.Router, validate *validator.Validate) {
	t.db = db
	t.router = router
	t.validate = validate

	genericHandler := generic_handlers.New[types.TestGroup](
		db, router, validate, data.TESTGROUPS_COLLECTION, types.FilterableTestGroupProps(),
	)

	genericHandler.ImplementGenericCrud("testgroups")

	// Implement additional routes below
}
