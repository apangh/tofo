package cloudtrailutil

import (
	"context"

	"github.com/apangh/tofo/cloudtrailutil/recorder"
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

type TrailRecorder struct {
	orm *model.ORM
}

func (r *TrailRecorder) Do(ctx context.Context, t types.Trail) error {
	o, e := recorder.ToTrail(&t)
	if e != nil {
		return e
	}
	return r.orm.TrailModel.Insert(ctx, o)
}
