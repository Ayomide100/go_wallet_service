package service

import (
	"context"
	"errors"
	"wallet-service/domain"
	"wallet-service/repository"
)

var ErrSameWallet = errors.New("cannot transfer to same wallet")

type WalletService struct {
	repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Transfer(
	ctx context.Context,
	fromID, toID string,
	amount domain.Money,
) error {

	if fromID == toID {
		return ErrSameWallet
	}

	from, err := s.repo.GetByID(ctx, fromID)
	if err != nil {
		return err
	}

	to, err := s.repo.GetByID(ctx, toID)
	if err != nil {
		return err
	}

	newFromBalance, err := from.Balance.Sub(amount)
	if err != nil {
		return err
	}

	from.Balance = newFromBalance
	to.Balance = to.Balance.Add(amount)

	if err := s.repo.Update(ctx, from); err != nil {
		return err
	}
	if err := s.repo.Update(ctx, to); err != nil {
		return err
	}

	return nil
}
