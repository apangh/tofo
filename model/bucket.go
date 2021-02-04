package model

import (
	"context"
	"errors"
	"time"
)

type Bucket struct {
	Name         string
	CreationDate time.Time
	Account      *Account
}

var (
	BucketAlreadyExist = errors.New("BucketAlreadyExists")
	BucketNotFound     = errors.New("BucketNowFound")
)

type BucketModel interface {
	Insert(ctx context.Context, b *Bucket) error
	Lookup(ctx context.Context, name string) (*Bucket, error)
	Dump(ctx context.Context)
}
