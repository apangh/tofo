package model

import (
	"context"
	"errors"
)

type User struct {
	DisplayName string
	Id          string
}

var (
	UserAlreadyExists = errors.New("UserAlreadyExists")
	UserNotFound      = errors.New("UserNotFound")
)

type UserModel interface {
	Insert(ctx context.Context, Id, displayName string) (*User, error)
	Lookup(ctx context.Context, Id string) (*User, error)
}
