package repository

import (
	"context"
	"errors"
	"sync"
	"wallet-service/domain"
)

type InMemoryWalletRepo struct {
	mu      sync.Mutex
	wallets map[string]*domain.Wallet
}

func NewInMemoryWalletRepo(wallets []*domain.Wallet) *InMemoryWalletRepo {
	m := make(map[string]*domain.Wallet)
	for _, w := range wallets {
		copy := *w
		m[w.ID] = &copy
	}
	return &InMemoryWalletRepo{wallets: m}
}

func (r *InMemoryWalletRepo) GetByID(ctx context.Context, id string) (*domain.Wallet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	w, ok := r.wallets[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	copy := *w
	return &copy, nil
}

func (r *InMemoryWalletRepo) Update(ctx context.Context, wallet *domain.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.wallets[wallet.ID] = wallet
	return nil
}
