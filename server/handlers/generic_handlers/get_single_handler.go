package generic_handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	hUtils "github.com/unexpectedtokens/api-tester/server/handlers/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *GenericHandler[T]) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	reqLog := httplog.LogEntry(r.Context())

	id := chi.URLParam(r, "id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	document, err := g.DAO.GetDocument(g.Collection, bson.M{"_id": objectId}, r.Context())
	if err != nil {

		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
		} else {
			reqLog.Info().Msg(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	hUtils.SendJSON(document, w)
}
