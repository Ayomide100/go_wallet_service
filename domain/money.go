package domain

import "errors"

var ErrInvalidAmount = errors.New("amount must be greater than zero")

type Money struct {
	cents int64
}

func NewMoneyFromCents(cents int64) (Money, error) {
	if cents <= 0 {
		return Money{}, ErrInvalidAmount
	}
	return Money{cents: cents}, nil
}

func (m Money) Cents() int64 {
	return m.cents
}

func (m Money) Add(other Money) Money {
	return Money{cents: m.cents + other.cents}
}

func (m Money) Sub(other Money) (Money, error) {
	if m.cents < other.cents {
		return Money{}, errors.New("insufficient funds")
	}
	return Money{cents: m.cents - other.cents}, nil
}
