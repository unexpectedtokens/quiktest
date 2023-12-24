package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	testcasehandlers "github.com/unexpectedtokens/api-tester/server/handlers/testcase_handlers"
	testgrouphandlers "github.com/unexpectedtokens/api-tester/server/handlers/testgroup_handlers"
	testreporthandlers "github.com/unexpectedtokens/api-tester/server/handlers/testreport_handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler interface {
	Initialize(db *mongo.Database, router chi.Router, validate *validator.Validate)
}

// TODO: Privatize
type Register struct {
	Validate *validator.Validate
	Router   chi.Router
	database *mongo.Database
}

func NewRegister(db *mongo.Database, router chi.Router, validate *validator.Validate) *Register {
	return &Register{database: db, Router: router, Validate: validate}
}

func (rs *Register) RegisterHandlers() {

	routeToHandlerMapping := &[]Handler{
		&testcasehandlers.TestCaseHandler{},
		&testgrouphandlers.TestGroupHandler{},
		&testreporthandlers.TestReportHandler{},
	}

	// Register api endpoints
	rs.Router.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))

		for _, handler := range *routeToHandlerMapping {
			handler.Initialize(rs.database, r, rs.Validate)
		}

	})
}
