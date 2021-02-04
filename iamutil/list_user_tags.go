package iamutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type TagCB interface {
	Do(ctx context.Context, t types.Tag) error
}

func ListUserTags(ctx context.Context, client *iam.Client, name string, cb TagCB) error {
	var marker *string
	for {
		params := &iam.ListUserTagsInput{
			Marker:   marker,
			UserName: aws.String(name),
			MaxItems: aws.Int32(100),
		}
		o, e := client.ListUserTags(ctx, params)
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
