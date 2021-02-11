package cloudtrailutil

import (
	"context"
	"fmt"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

type TrailStatusCB struct {
	client *cloudtrail.Client
}

func (ts *TrailStatusCB) Do(ctx context.Context, t *model.Trail) error {
	params := cloudtrail.GetTrailStatusInput{
		Name: aws.String(t.Name),
	}
	o, e := ts.client.GetTrailStatus(ctx, &params)
	if e != nil {
		return e
	}
	fmt.Printf("TrailStatus[%s]: %+v\n", t.Name, o)
	return nil
}

func Walk(ctx context.Context, client *cloudtrail.Client, orm *model.ORM) error {
	if e := DescribeTrails(ctx, client,
		&TrailRecorder{orm: orm}); e != nil {
		return e
	}
	if e := orm.TrailModel.Scan(ctx, &TrailStatusCB{client: client}); e != nil {
		return e
	}
	return nil
}
