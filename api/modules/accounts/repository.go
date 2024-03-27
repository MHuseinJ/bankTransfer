package accounts

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AccountRepo struct {
	httpClient http.Client
}

func NewRepo() AccountRepo {
	return AccountRepo{}
}

func (tr AccountRepo) validateUser(bankAccount *Account) (bool, error) {
	var client = &http.Client{}
	var data BankAccountStatus

	bodyRequest, err := json.Marshal(bankAccount)
	var payload = bytes.NewBufferString(string(bodyRequest))

	request, err := http.NewRequest("POST", "http://localhost:8002/bank/validate", payload)
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

	return data.Status, nil
}
