package accounts

import (
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	service *AccountService
}

func NewHandler(service *AccountService) AccountHandler {
	return AccountHandler{service}
}

func (th AccountHandler) Validate(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var account BankAccount
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"error": "Error Decode Request"}`))
		return
	}
	foundAccount, err := th.service.validate(account.AccountNumber)
	if err != nil {
		response.WriteHeader(404)
		response.Write([]byte(`{"error": "Account Not Found"}`))
		return
	}
	if foundAccount.Name != account.Name {
		response.WriteHeader(404)
		response.Write([]byte(`{"error": "Account Name Is Not Same"}`))
		return
	} else {
		respBody := BankAccountStatus{
			Name:          foundAccount.Name,
			Status:        foundAccount.Status,
			AccountNumber: foundAccount.AccountNumber,
		}
		resJson, err := json.Marshal(respBody)
		if err != nil {
			response.WriteHeader(500)
			response.Write([]byte(`{"error": "Error Decode Request"}`))
			return
		}
		response.Header().Add("Content-Type", "application/json")
		response.WriteHeader(200)
		response.Write(resJson)
	}
}
