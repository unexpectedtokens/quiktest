package generic_handlers

import (
	"net/http"

	"github.com/go-chi/httplog"
	hUtils "github.com/unexpectedtokens/api-tester/server/handlers/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *GenericHandler[T]) GetListHandler(w http.ResponseWriter, r *http.Request) {
	reqLog := httplog.LogEntry(r.Context())

	mongoQuery := bson.M{}

	urlQuery := r.URL.Query()
	if g.FilterableProps != nil {
		for _, prop := range *g.FilterableProps {
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
	}
	documents, err := g.DAO.GetDocuments(g.Collection, mongoQuery, r.Context())
	if err != nil {
		reqLog.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hUtils.SendJSON(documents, w)
}
