package s3util

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BucketRecorder struct {
	orm     *model.ORM
	account *model.Account
}

func (r *BucketRecorder) toBucket(b types.Bucket) *model.Bucket {
	return &model.Bucket{
		Name:         aws.ToString(b.Name),
		CreationDate: *b.CreationDate,
		Account:      r.account,
	}
}

func (r *BucketRecorder) Do(ctx context.Context, bucket types.Bucket) (
	*model.Bucket, error) {
	b := r.toBucket(bucket)
	if e := r.orm.BucketModel.Insert(ctx, b); e != nil {
		return nil, e
	}
	return b, nil
}

func toAccount(o types.Owner) *model.Account {
	return &model.Account{
		Id:   aws.ToString(o.ID),
		Name: aws.ToString(o.DisplayName),
	}
}
