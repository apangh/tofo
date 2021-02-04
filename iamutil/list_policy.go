package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type PolicyCB interface {
	Do(ctx context.Context, p types.Policy) error
}

func ListPolicies(ctx context.Context, client *iam.Client, cb PolicyCB) error {
	var marker *string
	for {
		params := &iam.ListPoliciesInput{
			Marker:       marker,
			MaxItems:     aws.Int32(100),
			OnlyAttached: true,
		}
		o, e := client.ListPolicies(ctx, params)
		if e != nil {
			return e
		}
		for _, p := range o.Policies {
			if e := cb.Do(ctx, p); e != nil {
				return e
			}
		}
		if !o.IsTruncated {
			return nil
		}
		marker = o.Marker
	}
}
