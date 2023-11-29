package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateDocumentHandler[T types.DocumentModel](db *mongo.Database, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())

		id := chi.URLParam(r, "id")

		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			reqLog.Warn().Msgf("error parsing objectId %s: %s", id, err.Error())
		}
		var respBody T

		reqLog.Info().Msg(fmt.Sprintf("attempting update of %s document", respBody.GetCollectionName()))

		err = json.NewDecoder(r.Body).Decode(&respBody)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			reqLog.Warn().Msgf("error decoding json body: %s", err.Error())
			return
		}

		err = validate.Struct(respBody)

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			reqLog.Warn().Msgf("validation error: %s", err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		reqLog.Info().Msgf("Updating '%s' document with id %s\n", respBody.GetCollectionName(), id)
		err = data.UpdateDocumentById(db, r.Context(), objectId, respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reqLog.Warn().Msg(fmt.Errorf("error updating document: %w", err).Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
