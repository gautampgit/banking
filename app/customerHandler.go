package app

import (
	"encoding/json"
	"net/http"

	"github.com/gautampgit/banking/service"
	"github.com/gorilla/mux"
)

type CustomHandler struct {
	service service.CustomerService
}

func (ch *CustomHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		//log.Fatal("Cannot find the customers")
	}
	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomerById(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
