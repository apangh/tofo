package iamutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func Walk(ctx context.Context, client *iam.Client, orm *model.ORM) error {
	if e := GetAccountAuthorizationDetails(ctx, client,
		&GroupDetailRecorder{orm: orm},
		&ManagedPolicyDetailRecorder{orm: orm},
		&RoleDetailRecorder{orm: orm},
		&UserDetailRecorder{orm: orm}); e != nil {
		return e
	}
	return nil
}
