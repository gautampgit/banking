package app

import (
	"encoding/json"
	"net/http"

	"github.com/gautampgit/banking/dto"

	"github.com/gorilla/mux"

	"github.com/gautampgit/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, apperr := h.service.NewAccount(request)
		if apperr != nil {
			writeResponse(w, apperr.Code, apperr.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	//decode incoming request
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	request.AccountId = accountId
	request.CustomerId = customerId

	//make the transaction
	account, apperr := h.service.MakeTransaction(request)
	if apperr != nil {
		writeResponse(w, apperr.Code, apperr.Message)
	}
	writeResponse(w, http.StatusOK, account)
}
