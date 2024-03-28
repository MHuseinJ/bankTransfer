package transactions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type TransactionService struct {
}

func NewService() TransactionService {
	return TransactionService{}
}

func (ts TransactionService) setlement(transaction Transaction) (bool, error) {
	var client = &http.Client{}
	var data Transaction

	bodyRequest, err := json.Marshal(transaction)
	var payload = bytes.NewBufferString(string(bodyRequest))

	request, err := http.NewRequest("POST", os.Getenv("API_HOST")+"/disbursement-callback", payload)
	if err != nil {
		return false, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ts TransactionService) initTransfer(transaction Transaction) (Transaction, error) {
	newTransaction := Transaction{
		Id:                 transaction.Id,
		Amount:             transaction.Amount,
		ReferenceID:        transaction.ReferenceID,
		OriginAccount:      transaction.OriginAccount,
		DestinationAccount: transaction.DestinationAccount,
		Status:             "HOLD",
	}
	return newTransaction, nil
}
