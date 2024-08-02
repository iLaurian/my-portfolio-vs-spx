package service

import "github.com/iLaurian/my-portfolio-vs-spx/entity"

type TransactionService interface {
	Add(entity.Transaction) entity.Transaction
	Edit(entity.Transaction) entity.Transaction
	Delete(entity.Transaction) entity.Transaction
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

func (service *transactionService) Edit(transaction entity.Transaction) entity.Transaction {
	for i, t := range service.transactions {
		if t.ID == transaction.ID {
			service.transactions[i] = transaction
			return transaction
		}
	}
	return entity.Transaction{}
}

func (service *transactionService) Delete(transaction entity.Transaction) entity.Transaction {
	for i, t := range service.transactions {
		if t.ID == transaction.ID {
			service.transactions = append(service.transactions[:i], service.transactions[i+1:]...)
			return transaction
		}
	}
	return entity.Transaction{}
}
