package model

import (
	"context"
	"errors"
	"time"
)

type Bucket struct {
	Name         string
	CreationDate time.Time
	Owner        *User
}

var (
	BucketAlreadyExist = errors.New("BucketAlreadyExists")
	BucketNotFound     = errors.New("BucketNowFound")
)

type BucketModel interface {
	Insert(ctx context.Context, name string, creationDate time.Time,
		owner *User) (*Bucket, error)
	Lookup(ctx context.Context, name string) (*Bucket, error)
}
