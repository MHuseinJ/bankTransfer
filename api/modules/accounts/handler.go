package accounts

import (
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	service *AccountService
}

func NewHandler(service *AccountService)  AccountHandler{
	return AccountHandler{service}
}

func(ah AccountHandler) Validate(response http.ResponseWriter, request *http.Request)  {
	response.Header().Add("Content-Type", "application/json")
	var accBank Account
	err := json.NewDecoder(request.Body).Decode(&accBank)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"error": "Error Decode Request"}`))
		return
	}
	status := ah.service.validateAccount(accBank)
	var bankResponse = BankAccountStatus{
		AccountNumber: accBank.AccountNumber,
		Status:        status,
		Name:          accBank.Name,
	}
	json.NewEncoder(response).Encode(&bankResponse)
}