package html_handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func testReportDetailPage(db *mongo.Database, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		testReport, err := data.GetDocument[types.TestReport](db, data.TESTREPORTS_COLLECTION, bson.M{"_id": objectId}, r.Context())
		if err != nil {

			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		testcaseResults, err := data.GetDocuments[types.TestCaseResult](db, data.TESTCASERESULTS_COLLECTION, bson.M{"testReportId": objectId}, r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		pageData := testReportDetailpageData{
			TestReport:      testReport,
			TestcaseResults: testcaseResults,
		}

		err = tmpl.ExecuteTemplate(w, indexTemplate, pageData)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func testReportsPage(db *mongo.Database, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		testreports, err := data.GetDocuments[types.TestReport](db, data.TESTREPORTS_COLLECTION, bson.M{}, r.Context())

		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		testCasePageData := testReportsPageData{
			TestReports: testreports,
		}

		err = tmpl.ExecuteTemplate(w, indexTemplate, testCasePageData)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
