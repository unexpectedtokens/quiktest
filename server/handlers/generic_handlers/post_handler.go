package generic_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/httplog"
	types "github.com/unexpectedtokens/api-tester/common_types"
	hUtils "github.com/unexpectedtokens/api-tester/server/handlers/util"
)

func (g *GenericHandler[T]) PostDocumentHandler(w http.ResponseWriter, r *http.Request) {
	reqLog := httplog.LogEntry(r.Context())
	var respBody T

	reqLog.Info().Msg(fmt.Sprintf("attempting creation of %s document", g.Collection))

	err := json.NewDecoder(r.Body).Decode(&respBody)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		reqLog.Warn().Msgf("error decoding json body: %s", err.Error())
		return
	}

	err = g.Validate.Struct(respBody)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		reqLog.Warn().Msgf("validation error: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	reqLog.Info().Msgf("Inserting document into collection '%s'\n", g.Collection)
	id, err := g.DAO.SaveDocument(respBody, r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		reqLog.Warn().Msg(fmt.Errorf("error saving document: %w", err).Error())
		return
	}

	hUtils.SendJSON(types.CreatedIdResponse{
		ID: id,
	}, w)
}
