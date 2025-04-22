package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplitedAPIKey

	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err

	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	account.AddBalancelance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) ListAccounts() ([]dto.AccountOutput, error) {
	accounts, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	output := make([]dto.AccountOutput, 0, len(accounts))
	for _, acc := range accounts {
		output = append(output, dto.AccountOutput{
			ID:      acc.ID,
			Name:    acc.Name,
			Email:   acc.Email,
			Balance: acc.Balance,
		})
	}

	return output, nil
}
