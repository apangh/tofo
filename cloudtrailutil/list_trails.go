package cloudtrailutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

type TrailInfoCB interface {
	Do(ctx context.Context, trail types.TrailInfo) error
}

func ListTrails(ctx context.Context, client *cloudtrail.Client, cb TrailInfoCB) error {
	params := &cloudtrail.ListTrailsInput{}
	p := cloudtrail.NewListTrailsPaginator(client, params)
	for p.HasMorePages() {
		o, e := p.NextPage(ctx)
		if e != nil {
			return e
		}
		for _, trail := range o.Trails {
			if e := cb.Do(ctx, trail); e != nil {
				return e
			}
		}
	}
	return nil
}
