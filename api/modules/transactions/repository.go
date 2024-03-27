package transactions

import (
	"bytes"
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) TransactionRepo {
	return TransactionRepo{DB}
}

func (tr TransactionRepo) saveTransaction(transaction Transaction) (Transaction, error) {
	result := tr.DB.Create(&transaction)
	return transaction, result.Error
}

func (tr TransactionRepo) updateTransaction(transaction Transaction, transaction2 Transaction) (Transaction, error) {
	result := tr.DB.Model(&transaction).Updates(transaction2)
	return transaction, result.Error
}

func (tr TransactionRepo) findTransactionByReferenceId(refId string) (Transaction, error) {
	var transaction Transaction
	result := tr.DB.First(&transaction, "reference_id = ?", refId)
	return transaction, result.Error
}

func (tr TransactionRepo) callTransferBank(transaction Transaction) (Transaction, error) {
	var client = &http.Client{}
	var data Transaction

	bodyRequest, err := json.Marshal(transaction)
	var payload = bytes.NewBufferString(string(bodyRequest))

	request, err := http.NewRequest("POST", "http://localhost:8002/bank/transfer", payload)
	if err != nil {
		return Transaction{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return Transaction{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return Transaction{}, err
	}

	return data, nil
}
