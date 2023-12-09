package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/go-playground/validator/v10"
	types "github.com/unexpectedtokens/api-tester/common_types"
	"github.com/unexpectedtokens/api-tester/server/data"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostDocumentHandler[T types.DocumentModel](db *mongo.Database, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqLog := httplog.LogEntry(r.Context())
		var respBody T

		reqLog.Info().Msg(fmt.Sprintf("attempting creation of %s document", respBody.GetCollectionName()))

		err := json.NewDecoder(r.Body).Decode(&respBody)

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

		reqLog.Info().Msgf("Inserting document into collection '%s'\n", respBody.GetCollectionName())
		id, err := data.SaveDocument(db, respBody, r.Context(), reqLog)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reqLog.Warn().Msg(fmt.Errorf("error saving document: %w", err).Error())
			return
		}

		fmt.Println(id)
		jsonSend[types.CreatedIdResponse](types.CreatedIdResponse{
			ID: id,
		}, w)
	}
}
