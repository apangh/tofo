package iamutil

import (
	"context"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type GroupDetailCB interface {
	Do(ctx context.Context, groupD types.GroupDetail) error
}

type RoleDetailCB interface {
	Do(ctx context.Context, roleD types.RoleDetail) error
}

type PolicyDetailCB interface {
	Do(ctx context.Context, policyD types.ManagedPolicyDetail) error
}

type UserDetailCB interface {
	Do(ctx context.Context, userD types.UserDetail) error
}

func GetAccountAuthorizationDetails(ctx context.Context, client *iam.Client,
	groupDetailCB GroupDetailCB, policyDetailCB PolicyDetailCB,
	roleDetailCB RoleDetailCB, userDetailCB UserDetailCB) error {
	params := &iam.GetAccountAuthorizationDetailsInput{}
	p := iam.NewGetAccountAuthorizationDetailsPaginator(client, params)
	for p.HasMorePages() {
		o, e := p.NextPage(ctx)
		if e != nil {
			return e
		}
		if !reflect.ValueOf(groupDetailCB).IsNil() {
			for _, groupD := range o.GroupDetailList {
				if e := groupDetailCB.Do(ctx, groupD); e != nil {
					return e
				}
			}
		}
		if !reflect.ValueOf(policyDetailCB).IsNil() {
			for _, policyD := range o.Policies {
				if e := policyDetailCB.Do(ctx, policyD); e != nil {
					return e
				}
			}
		}
		if !reflect.ValueOf(roleDetailCB).IsNil() {
			for _, roleD := range o.RoleDetailList {
				if e := roleDetailCB.Do(ctx, roleD); e != nil {
					return e
				}
			}
		}
		if !reflect.ValueOf(userDetailCB).IsNil() {
			for _, userD := range o.UserDetailList {
				if e := userDetailCB.Do(ctx, userD); e != nil {
					return e
				}
			}
		}
	}
	return nil
}
