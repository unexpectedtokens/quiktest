package testreporthandlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"github.com/unexpectedtokens/api-tester/server/handlers/generic_handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestReportHandler struct {
	db       *mongo.Database
	router   chi.Router
	validate *validator.Validate
}

func (t *TestReportHandler) Initialize(db *mongo.Database, router chi.Router, validate *validator.Validate) {
	t.db = db
	t.router = router
	t.validate = validate

	genericReportsCrudHandler := generic_handlers.New[types.TestReport](
		db, router, validate, data.TESTREPORTS_COLLECTION, nil,
	)
	genericCaseResultCrudHandler := generic_handlers.New[types.TestCaseResult](
		db, router, validate, data.TESTCASERESULTS_COLLECTION, types.FilterableTestcaseProps(),
	)

	t.router.Get("/testreports", genericReportsCrudHandler.GetListHandler)
	t.router.Get("/testreports/{id}", genericReportsCrudHandler.GetByIDHandler)
	t.router.Post("/testreports", genericReportsCrudHandler.PostDocumentHandler)

	t.router.Get("/testreports/results", genericCaseResultCrudHandler.GetListHandler)
	t.router.Post("/testreports/results", genericCaseResultCrudHandler.PostDocumentHandler)

	// Implement additional routes below

}
