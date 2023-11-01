package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func jsonSend[T any](res T, rw http.ResponseWriter) {
	binResponse, err := json.Marshal(res)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
	rw.Write(binResponse)
}

func getObjectIdFromRequest(r *http.Request, param string) (primitive.ObjectID, error) {
	id := chi.URLParam(r, "id")
	reqLog := httplog.LogEntry(r.Context())

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		reqLog.Error().Msgf("error parsing objectId '%s': %s", id, err.Error())
		return objectId, err
	}

	return objectId, nil
}
