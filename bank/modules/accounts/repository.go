package accounts

import "gorm.io/gorm"

type AccountRepo struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) AccountRepo {
	return AccountRepo{DB}
}

func (tr AccountRepo) findAccountByAccountNumber(accountNumber string) (Account, error) {
	var account Account
	result := tr.DB.First(&account, "account_number=?", accountNumber)
	return account, result.Error
}
