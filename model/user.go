package model

import (
	"context"
	"errors"
	"time"
)

type User struct {
	Id                 string
	Name               string
	Path               string
	Arn                string
	CreateDate         *time.Time
	PasswordLastUsed   *time.Time
	Tags               map[string]string
	PermissionBoundary string // TODO
}

var (
	UserAlreadyExists = errors.New("UserAlreadyExists")
	UserNotFound      = errors.New("UserNotFound")
)

type UserModel interface {
	Insert(ctx context.Context, u *User) error
	Lookup(ctx context.Context, Id string) (*User, error)
	Dump(ctx context.Context)
}