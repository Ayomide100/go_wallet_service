package repository

import (
	"context"
	"wallet-service/domain"
)

type WalletRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Wallet, error)
	Update(ctx context.Context, wallet *domain.Wallet) error
}
