package recorder

import (
	"github.com/apangh/tofo/model"
	commonRecorder "github.com/apangh/tofo/recorder"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToAttachedPermissionsBoundary(b *types.AttachedPermissionsBoundary) (
	*model.AttachedPermissionsBoundary, error) {
	if b == nil {
		return nil, nil
	}
	arn, e := commonRecorder.ToArn(b.PermissionsBoundaryArn)
	if e != nil {
		return nil, e
	}
	return &model.AttachedPermissionsBoundary{
		Arn: arn,
	}, nil
}
