package iamutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func Walk(ctx context.Context, client *iam.Client, orm *model.ORM) error {
	if e := ListPolicies(ctx, client, &PolicyRecorder{orm: orm}); e != nil {
		return e
	}
	if e := ListUsers(ctx, client, &UserRecorder{orm: orm}); e != nil {
		return e
	}
	if e := ListRoles(ctx, client, &RoleRecorder{orm: orm}); e != nil {
		return e
	}
	if e := ListGroups(ctx, client, &GroupRecorder{orm: orm}); e != nil {
		return e
	}
	return nil
}
