package repository

import (
	"context"
	"errors"
	"wallet-service/domain"
)

type FailingWalletRepo struct{}

func (f *FailingWalletRepo) GetByID(ctx context.Context, id string) (*domain.Wallet, error) {
	return nil, errors.New("repository unavailable")
}

func (f *FailingWalletRepo) Update(ctx context.Context, wallet *domain.Wallet) error {
	return errors.New("repository unavailable")
}
