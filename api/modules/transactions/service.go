package transactions

import "log"

type TransactionService struct {
	repo *TransactionRepo
}

func NewService(repo *TransactionRepo) TransactionService {
	return TransactionService{repo}
}

func (ts TransactionService) makeTransfer(transaction Transaction) bool {
	transaction, err := ts.repo.callTransferBank(transaction)
	if err != nil {
		log.Fatalf(err.Error())
		return false
	}
	ts.repo.saveTransaction(transaction)
	return true
}

func (ts TransactionService) updateTransferStatus(transaction Transaction) (Transaction, error) {
	transactionToUpdate, err := ts.repo.findTransactionByReferenceId(transaction.ReferenceID)
	if err != nil {
		return Transaction{}, err
	}
	updatedStatusTransaction := Transaction{
		Status:      transaction.Status,
		ReferenceID: transactionToUpdate.ReferenceID,
	}
	transactionUpdated, err := ts.repo.updateTransaction(transactionToUpdate, updatedStatusTransaction)
	if err != nil {
		return Transaction{}, err
	}
	return transactionUpdated, nil
}
