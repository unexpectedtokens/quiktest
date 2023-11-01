package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	types "github.com/unexpectedtokens/api-tester/common"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetListHandler[T types.DocumentModel](db *mongo.Database, collection string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())

		documents, err := data.GetDocuments[T](db, collection, bson.M{}, r.Context())
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
