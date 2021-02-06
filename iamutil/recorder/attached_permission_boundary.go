package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToAttachedPermissionsBoundary(b *types.AttachedPermissionsBoundary) *model.AttachedPermissionsBoundary {
	if b == nil {
		return nil
	}
	return &model.AttachedPermissionsBoundary{
		Arn: b.PermissionsBoundaryArn,
	}
}
