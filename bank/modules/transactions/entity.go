package transactions

type Transaction struct {
	Id                 int64  `gorm:"primary key;autoIncrement" json:"id"`
	OriginAccount      string `json:"originAccount"`
	Amount             int64  `json:"amount"`
	DestinationAccount string `json:"destinationAccount"`
	ReferenceID        string `gorm:"unique" json:"referenceID"`
	Status             string `gorm:"default:HOLD" json:"status"`
}
