package app

import (
	"encoding/json"
	"net/http"

	"github.com/gautampgit/banking/dto"

	"github.com/gautampgit/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	account, Apperr := h.service.NewAccount(request)
	if Apperr != nil {
		writeResponse(w, Apperr.Code, Apperr.Message)
	}
	writeResponse(w, http.StatusCreated, account)
}
