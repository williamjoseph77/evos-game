package handlers

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/williamjoseph77/evos/objects"
	"github.com/williamjoseph77/evos/services"
)

func (h *Handler) HandleGetRoleList(responseWriter http.ResponseWriter, request *http.Request) {
	var responsePayload []objects.GetRoleListResponse
	roles, err := services.GetRoleList(h.Database)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	err = copier.Copy(&responsePayload, &roles)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}
