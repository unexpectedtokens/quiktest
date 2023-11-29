package html_handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func testCasesPage(db *mongo.Database, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		testcases, err := data.GetDocuments[types.TestCase](db, data.TESTCASES_COLLECTION, bson.M{}, r.Context())

		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		testCasePageData := testCasePageData{
			Testcases: testcases,
		}

		err = tmpl.ExecuteTemplate(w, indexTemplate, testCasePageData)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func newTestCasePage(db *mongo.Database, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, indexTemplate, nil)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func handleCreateNewTestcase(db *mongo.Database, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for key, values := range r.Form { // range over map
			for _, value := range values { // range over []string
				fmt.Println(key, value)
			}
		}

		retCode, err := strconv.Atoi(r.FormValue("expectReturnCode"))

		if err != nil {
			log.Println(err)
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		newTestcase := types.TestCase{
			Title:            r.FormValue("title"),
			ExpectReturnCode: retCode,
			Request: types.TestRequest{
				Route:  r.FormValue("route"),
				Method: r.FormValue("method"),
			},
		}

		err = validate.Struct(newTestcase)

		if err != nil {
			log.Println(err)
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/testcases", http.StatusFound)
	}
}
