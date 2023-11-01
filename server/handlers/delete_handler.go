package handlers

import (
	"net/http"

	"github.com/go-chi/httplog"
	types "github.com/unexpectedtokens/api-tester/common"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteDocumentByIDHandler[T types.DocumentModel](db *mongo.Database, collection string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())
		objectId, err := getObjectIdFromRequest(r, "id")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = data.DeleteDocument[T](db, collection, r.Context(), bson.M{"_id": objectId})
		if err != nil {
			reqLog.Warn().Msgf("error deleting document with id %s: %s", objectId.Hex(), err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reqLog.Info().Msgf("deleted document with id %s", objectId.Hex())
		w.WriteHeader(http.StatusNoContent)
	}
}
