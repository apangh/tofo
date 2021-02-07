package model

import (
	"context"
	"errors"
	"time"
)

type Group struct {
	Id         string
	Name       string
	Path       string
	Arn        *ARN
	CreateDate time.Time
}

var (
	GroupAlreadyExists = errors.New("GroupAlreadyExists")
	GroupNotFound      = errors.New("GroupNotFound")
)

type GroupModel interface {
	Insert(ctx context.Context, g *Group) error
	Lookup(ctx context.Context, id string) (*Group, error)
	Dump(ctx context.Context)
}
