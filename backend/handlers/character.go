package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/williamjoseph77/evos/objects"
	"github.com/williamjoseph77/evos/services"
)

func (h *Handler) HandleCreateCharacterNonSecure(responseWriter http.ResponseWriter, request *http.Request) {
	var requestPayload objects.CreateCharacterNonSecureRequest
	var responsePayload objects.CreateCharacterResponse

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&requestPayload); err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer request.Body.Close()

	character, err := services.CreateCharacterNonSecure(h.Database, requestPayload)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	responsePayload.GUID = character.GUID

	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}

func (h *Handler) HandleCreateCharacterSecure(responseWriter http.ResponseWriter, request *http.Request) {
	var requestPayload objects.CreateCharacterSecureRequest
	var responsePayload objects.CreateCharacterResponse

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&requestPayload); err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	character, err := services.CreateCharacterSecure(h.Database, requestPayload)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	responsePayload.GUID = character.GUID

	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}

func (h *Handler) HandleGetCharacterList(responseWriter http.ResponseWriter, request *http.Request) {
	var responsePayload []objects.GetCharacterListResponse
	characters, err := services.GetCharacterList(h.Database)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	err = copier.Copy(&responsePayload, &characters)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}
