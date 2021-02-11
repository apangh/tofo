package cloudtrailutil

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func Walk(ctx context.Context, client *cloudtrail.Client, orm *model.ORM) error {
	if e := DescribeTrails(ctx, client,
		&TrailRecorder{orm: orm}); e != nil {
		return e
	}
	return nil
}
