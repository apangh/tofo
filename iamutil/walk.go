package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func Walk(ctx context.Context, client *iam.Client) error {
	var cb LogUser
	return ListUsers(ctx, client, &cb)
}
