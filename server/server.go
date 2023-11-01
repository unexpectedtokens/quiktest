package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common"
	"github.com/unexpectedtokens/api-tester/server/data"
	"github.com/unexpectedtokens/api-tester/server/db"
	"github.com/unexpectedtokens/api-tester/server/handlers"
)

func RunServer() {
	logger := httplog.NewLogger("quiktest-api", httplog.Options{
		Concise: false,
	})

	validate := validator.New(validator.WithRequiredStructEnabled())

	connection, err := db.NewConnection()

	if err != nil {
		panic(err)
	}

	defer connection.Disconnect(context.Background())

	db := db.GetDB(connection)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Register middleware
	r.Use(httplog.RequestLogger(logger, []string{"/ping", "/public"}))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer, middleware.RequestID)

	// Register api endpoints
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		// Testcases
		r.Get("/testcases", handlers.GetListHandler[types.TestCase](db, data.TESTCASES_COLLECTION))

		r.Post("/testcases", handlers.PostDocumentHandler[types.TestCase](db, validate))

		r.Delete("/testcases/{id}", handlers.DeleteDocumentByIDHandler[types.TestCase](db, data.TESTCASES_COLLECTION))

		r.Put("/testcases/{id}", handlers.UpdateDocumentHandler[types.TestCase](db, validate))

		// Testreports
		r.Get("/testreports", handlers.GetListHandler[types.TestReport](db, data.TESTREPORTS_COLLECTION))

		r.Post("/testreports", handlers.PostDocumentHandler[types.TestReport](db, validate))

		r.Get("/testreports/{id}/results", handlers.GetTestCaseResultsByReportId(db))

		r.Post("/testreports/results", handlers.PostDocumentHandler[types.TestCaseResult](db, validate))

		r.Get("/testreports/{id}", handlers.GetByIDHandler[types.TestReport](db, data.TESTREPORTS_COLLECTION))

	})

	logger.Info().Msg("Registered the following routes")

	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	})

	logger.Info().Msg("Running on 8080")
	http.ListenAndServe("localhost:8080", r)
}
