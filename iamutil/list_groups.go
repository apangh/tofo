package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type GroupCB interface {
	Do(ctx context.Context, group types.Group) error
}

func ListGroups(ctx context.Context, client *iam.Client, cb GroupCB) error {
	params := &iam.ListGroupsInput{
		MaxItems: aws.Int32(100),
	}
	p := iam.NewListGroupsPaginator(client, params)
	for p.HasMorePages() {
		o, e := p.NextPage(ctx)
		if e != nil {
			return e
		}
		for _, group := range o.Groups {
			if e := cb.Do(ctx, group); e != nil {
				return e
			}
		}
	}
	return nil
}
