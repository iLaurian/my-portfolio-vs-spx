package service

import (
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/repository"
)

type TransactionService interface {
	Add(entity.Transaction) error
	Edit(entity.Transaction) error
	Delete(int) error
	FindAll() ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func New(repository repository.TransactionRepository) TransactionService {
	return &transactionService{
		repository: repository,
	}
}

func (s *transactionService) FindAll() ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAll()
	return transactions, err
}

func (s *transactionService) Add(transaction entity.Transaction) error {
	err := s.repository.Add(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) Edit(transaction entity.Transaction) error {
	err := s.repository.Edit(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
