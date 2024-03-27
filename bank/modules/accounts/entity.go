package accounts

type Account struct {
	Id            int64  `gorm:"primary key;autoIncrement" json:"id"`
	AccountNumber string `json:"accountNumber"`
	Name          string `json:"name"`
	Status        bool   `gorm:"default:true" json:"status"`
}

type BankAccount struct {
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
}

type BankAccountStatus struct {
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	Status        bool   `json:"status"`
}
