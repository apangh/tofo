package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func Walk(ctx context.Context, client *iam.Client) error {
	if e := ListUsers(ctx, client, &LogUser{}); e != nil {
		return e
	}
	if e := ListRoles(ctx, client, &LogRole{}); e != nil {
		return e
	}
	if e := ListGroups(ctx, client, &LogGroup{}); e != nil {
		return e
	}
	return nil
}
