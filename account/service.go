package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

// Service is an interface for Account methods
type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error)
}

// Account represents an user account
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}

	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	a, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *accountService) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	accounts, err := s.repository.ListAccounts(ctx, skip, take)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
