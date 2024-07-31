package service

import "github.com/iLaurian/my-portfolio-vs-spx/entity"

type TransactionService interface {
	Add(entity.Transaction) entity.Transaction
	FindAll() []entity.Transaction
}

type transactionService struct {
	transactions []entity.Transaction
}

func New() TransactionService {
	return &transactionService{}
}

func (service *transactionService) Add(transaction entity.Transaction) entity.Transaction {
	service.transactions = append(service.transactions, transaction)
	return transaction
}

func (service *transactionService) FindAll() []entity.Transaction {
	return service.transactions
}
