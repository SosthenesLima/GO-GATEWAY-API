package service

import "github.com/devfullcycle/imersao22/go-gateway/internal/domain"

type AccountService struct {
	repository domain.AcoountRepository
}

func NewAccountService(repository domain.AcoountRepository) *AccountService {
	return &AccountService{repository: repository}
}
