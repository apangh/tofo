package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListRoleTags(ctx context.Context, client *iam.Client, name string, cb TagCB) error {
	var marker *string
	for {
		params := &iam.ListRoleTagsInput{
			Marker:   marker,
			RoleName: aws.String(name),
			MaxItems: aws.Int32(100),
		}
		o, e := client.ListRoleTags(ctx, params)
		if e != nil {
			return e
		}
		for _, t := range o.Tags {
			if e := cb.Do(ctx, t); e != nil {
				return e
			}
		}

		if !o.IsTruncated {
			return nil
		}
		marker = o.Marker
	}
}
