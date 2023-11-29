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

func GetByIDHandler[T types.DocumentModel](db *mongo.Database, model T) http.HandlerFunc {
	collectionName := model.GetCollectionName()
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())

		id := chi.URLParam(r, "id")

		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		document, err := data.GetDocument[T](db, collectionName, bson.M{"_id": objectId}, r.Context())
		if err != nil {

			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
			} else {
				reqLog.Info().Msg(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		jsonSend[T](document, w)
	}
}
