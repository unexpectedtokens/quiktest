package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetByGroupHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// reqLog := httplog.LogEntry(r.Context())

		groupName := chi.URLParam(r, "groupname")

		fmt.Println(groupName)
		// testcases, err := data.GetDocuments[types.TestCaseResult](db, data.TESTCASERESULTS_COLLECTION, bson.M{"testReportId": objectId}, r.Context())

		// if err != nil {
		// 	reqLog.Warn().Msgf("error getting documents: %s", err.Error())
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// jsonSend[[]types.TestCaseResult](testCaseResults, w)
	}
}
