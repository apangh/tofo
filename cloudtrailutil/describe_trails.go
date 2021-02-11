package cloudtrailutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

type TrailCB interface {
	Do(ctx context.Context, trail types.Trail) error
}

func DescribeTrails(ctx context.Context, client *cloudtrail.Client, cb TrailCB) error {
	params := &cloudtrail.DescribeTrailsInput{}
	o, e := client.DescribeTrails(ctx, params)
	if e != nil {
		return e
	}
	for _, t := range o.TrailList {
		if e := cb.Do(ctx, t); e != nil {
			return e
		}
	}
	return nil
}
