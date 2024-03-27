package accounts

import "log"

type AccountService struct {
	repo *AccountRepo
}

func NewService(repo *AccountRepo) AccountService {
	return AccountService{repo}
}

func (as AccountService) validateAccount(bankAccount Account) bool {
	result, err := as.repo.validateUser(&bankAccount)
	if err != nil {
		log.Fatalf(err.Error())
		return false
	}
	return result
}
