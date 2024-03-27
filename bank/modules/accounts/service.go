package accounts

import "log"

type AccountService struct {
	repo *AccountRepo
}

func NewService(repo *AccountRepo) AccountService {
	return AccountService{repo}
}

func (as AccountService) validate(accountNumber string) (Account, error) {
	foundAccount, err := as.repo.findAccountByAccountNumber(accountNumber)
	if err != nil {
		log.Fatalf(err.Error())
		return Account{}, err
	}
	return foundAccount, nil
}
