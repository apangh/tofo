package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type RoleCB interface {
	Do(ctx context.Context, role types.Role) error
}

func ListRoles(ctx context.Context, client *iam.Client, cb RoleCB) error {
	params := &iam.ListRolesInput{
		MaxItems: aws.Int32(100),
	}
	p := iam.NewListRolesPaginator(client, params)
	for p.HasMorePages() {
		o, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, role := range o.Roles {
			if e := cb.Do(ctx, role); e != nil {
				return e
			}
		}
	}
	return nil
}
