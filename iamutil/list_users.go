package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type ListUsersCB interface {
	Do(ctx context.Context, user types.User) error
}

func ListUsers(ctx context.Context, client *iam.Client, cb ListUsersCB) error {
	params := &iam.ListUsersInput{
		MaxItems: aws.Int32(100),
	}
	p := iam.NewListUsersPaginator(client, params)
	for p.HasMorePages() {
		o, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, user := range o.Users {
			if e := cb.Do(ctx, user); e != nil {
				return e
			}
		}
	}
	return nil
}
