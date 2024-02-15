package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MarkTBSS/go-hexagonalArchitecture/service"
	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerService service.ICustomService
}

func NewCustomerHandler(customerService service.ICustomService) *customerHandler {
	return &customerHandler{
		customerService: customerService,
	}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.customerService.GetCustomer(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
