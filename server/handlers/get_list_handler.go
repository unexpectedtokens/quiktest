package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetListHandler[T types.FilterableDocumentModel](db *mongo.Database, model T) http.HandlerFunc {
	collectionName := model.GetCollectionName()
	filterableProps := model.FilterableProps()
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())

		mongoQuery := bson.M{}

		urlQuery := r.URL.Query()
		for _, prop := range *filterableProps {
			// TODO: Abstract
			if urlQuery.Has(prop) {
				val := urlQuery.Get(prop)

				reqLog.Info().Msgf("filter is %s\n\n\n", val)
				if val == "nil" {
					reqLog.Info().Msgf("filter is %s\n\n\n", val)

					mongoQuery[prop] = nil
				} else if id, err := primitive.ObjectIDFromHex(val); err == nil {
					mongoQuery[prop] = id
				} else {
					mongoQuery[prop] = urlQuery.Get(prop)
				}
			}
		}

		documents, err := data.GetDocuments[T](db, collectionName, mongoQuery, r.Context())
		if err != nil {
			reqLog.Info().Msg(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonSend[[]T](documents, w)
	}
}

func GetTestCaseResultsByReportId(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())

		id := chi.URLParam(r, "id")

		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			reqLog.Warn().Msgf("error parsing objectId %s: %s", id, err.Error())
			return
		}

		testCaseResults, err := data.GetDocuments[types.TestCaseResult](db, data.TESTCASERESULTS_COLLECTION, bson.M{"testReportId": objectId}, r.Context())

		if err != nil {
			reqLog.Warn().Msgf("error getting documents: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonSend[[]types.TestCaseResult](testCaseResults, w)
	}
}
