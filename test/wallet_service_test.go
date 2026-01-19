package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"wallet-service/domain"
	"wallet-service/repository"
	"wallet-service/service"
)

func TestSuccessfulTransfer(t *testing.T) {
	w1Bal, _ := domain.NewMoneyFromCents(1000)
	w2Bal, _ := domain.NewMoneyFromCents(500)

	repo := repository.NewInMemoryWalletRepo([]*domain.Wallet{
		{ID: "w1", Owner: "Alice", Balance: w1Bal},
		{ID: "w2", Owner: "Bob", Balance: w2Bal},
	})

	svc := service.NewWalletService(repo)
	amount, _ := domain.NewMoneyFromCents(300)

	err := svc.Transfer(context.Background(), "w1", "w2", amount)
	assert.NoError(t, err)

	w1, _ := repo.GetByID(context.Background(), "w1")
	w2, _ := repo.GetByID(context.Background(), "w2")

	assert.Equal(t, int64(700), w1.Balance.Cents())
	assert.Equal(t, int64(800), w2.Balance.Cents())
}

func TestInsufficientFunds(t *testing.T) {
	bal, _ := domain.NewMoneyFromCents(100)

	repo := repository.NewInMemoryWalletRepo([]*domain.Wallet{
		{ID: "w1", Owner: "Alice", Balance: bal},
		{ID: "w2", Owner: "Bob", Balance: bal},
	})

	svc := service.NewWalletService(repo)
	amount, _ := domain.NewMoneyFromCents(200)

	err := svc.Transfer(context.Background(), "w1", "w2", amount)
	assert.Error(t, err)
}

func TestFailingRepository(t *testing.T) {
	repo := &repository.FailingWalletRepo{}
	svc := service.NewWalletService(repo)

	amount, _ := domain.NewMoneyFromCents(100)

	err := svc.Transfer(context.Background(), "w1", "w2", amount)
	assert.Error(t, err)
}
