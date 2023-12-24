package generic_handlers

import (
	"net/http"

	"github.com/go-chi/httplog"
	hUtils "github.com/unexpectedtokens/api-tester/server/handlers/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (g *GenericHandler[T]) DeleteDocumentByIDHandler(w http.ResponseWriter, r *http.Request) {
	reqLog := httplog.LogEntry(r.Context())
	objectId, err := hUtils.ObjectIdFromRequest(r, "id")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = g.DAO.DeleteDocument(g.Collection, r.Context(), bson.M{"_id": objectId})
	if err != nil {
		reqLog.Warn().Msgf("error deleting document with id %s: %s", objectId.Hex(), err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqLog.Info().Msgf("deleted document with id %s", objectId.Hex())
	w.WriteHeader(http.StatusNoContent)
}
