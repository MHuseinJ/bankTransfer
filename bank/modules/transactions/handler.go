package transactions

import (
	"encoding/json"
	"log"
	"net/http"
)

type TransactionHandler struct {
	service *TransactionService
}

func NewHandler(service *TransactionService) TransactionHandler {
	return TransactionHandler{service}
}

func (th TransactionHandler) Settlement(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var transaction Transaction
	err := json.NewDecoder(request.Body).Decode(&transaction)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"error": "Error Decode Request"}`))
		return
	}
	_, err = th.service.setlement(transaction)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"error": "Error Call API Settlement"}`))
		log.Fatalf(err.Error())
		return
	}
	response.WriteHeader(200)
	response.Write([]byte(`{"status": "success", "message": "Transfer Complete"}`))
}

func (th TransactionHandler) InitTransfer(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var payload Transaction
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"error": "Error Decode Request"}`))
		return
	}
	updatedTransaction, err := th.service.initTransfer(payload)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`"status": "fail", "message": "Failed Updated Transaction"`))
		log.Fatalf(err.Error())
		return
	}
	response.WriteHeader(200)
	res, err := json.Marshal(updatedTransaction)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`"status": "fail", "message": "Failed Unmarshal response"`))
	}
	response.Write(res)
}
