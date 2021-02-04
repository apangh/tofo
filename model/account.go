package model

import (
	"context"
	"errors"
)

type Account struct {
	Id   string
	Name string
}

var (
	AccountAlreadyExists = errors.New("AccountAlreadyExists")
	AccountNotFound      = errors.New("AccountNotFound")
)

type AccountModel interface {
	Insert(ctx context.Context, a *Account) error
	Lookup(ctx context.Context, Id string) (*Account, error)
	Dump(ctx context.Context)
}
