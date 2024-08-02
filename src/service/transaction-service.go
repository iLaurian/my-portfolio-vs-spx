package service

import "github.com/iLaurian/my-portfolio-vs-spx/entity"

type TransactionService interface {
	Add(entity.Transaction) error
	Edit(entity.Transaction) error
	Delete(entity.Transaction) error
	FindAll() []entity.Transaction
}

type transactionService struct {
	transactions []entity.Transaction
}

func New() TransactionService {
	return &transactionService{}
}

func (service *transactionService) FindAll() []entity.Transaction {
	return service.transactions
}

func (service *transactionService) Add(transaction entity.Transaction) error {
	service.transactions = append(service.transactions, transaction)
	return nil
}

func (service *transactionService) Edit(transaction entity.Transaction) error {
	for i, t := range service.transactions {
		if t.ID == transaction.ID {
			service.transactions[i] = transaction
			return nil
		}
	}
	return nil
}

func (service *transactionService) Delete(transaction entity.Transaction) error {
	for i, t := range service.transactions {
		if t.ID == transaction.ID {
			service.transactions = append(service.transactions[:i], service.transactions[i+1:]...)
			return nil
		}
	}
	return nil
}
