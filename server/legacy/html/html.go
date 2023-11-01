package html_handlers

import (
	"embed"
	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	layout                  = "templates/index.html"
	testcases_page          = "templates/testcases.html"
	testcases_create_page   = "templates/newtestcase.html"
	testreports_page        = "templates/testreports.html"
	testreports_detail_page = "templates/testreportdetail.html"
	indexTemplate           = "index"
)

var (
	//go:embed templates/*
	files embed.FS
)

func InitHTMLRoutes(db *mongo.Database, validate *validator.Validate) (*chi.Mux, error) {
	r := chi.NewRouter()

	testCasesTemplate, err := template.ParseFS(files, layout, testcases_page)
	if err != nil {
		return nil, err
	}

	testReportsTemplate, err := template.ParseFS(files, layout, testreports_page)

	if err != nil {
		return nil, err
	}

	testReportDetailTemplate, err := template.ParseFS(files, layout, testreports_detail_page)

	if err != nil {
		return nil, err
	}

	testCaseCreatePage, err := template.ParseFS(files, layout, testcases_create_page)

	if err != nil {
		return nil, err
	}
	r.Get("/testcases", testCasesPage(db, testCasesTemplate))
	r.Post("/testcases", handleCreateNewTestcase(db, validate))
	r.Get("/testcases/new", newTestCasePage(db, testCaseCreatePage))
	r.Get("/testreports", testReportsPage(db, testReportsTemplate))
	r.Get("/testreports/{id}", testReportDetailPage(db, testReportDetailTemplate))
	return r, nil
}
