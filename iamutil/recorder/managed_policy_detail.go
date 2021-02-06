package recorder

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type ManagedPolicyDetailRecorder struct {
	Orm *model.ORM
}

func ToManagedPolicyDetail(p types.ManagedPolicyDetail) *model.ManagedPolicyDetail {
	return &model.ManagedPolicyDetail{
		Id:               aws.ToString(p.PolicyId),
		Name:             aws.ToString(p.PolicyName),
		Path:             aws.ToString(p.Path),
		Arn:              aws.ToString(p.Arn),
		AttachmentCount:  aws.ToInt32(p.AttachmentCount),
		DefaultVersionId: aws.ToString(p.DefaultVersionId),
		Description:      aws.ToString(p.Description),
		IsAttachable:     p.IsAttachable,
		CreateDate:       aws.ToTime(p.CreateDate),
		UpdateDate:       aws.ToTime(p.UpdateDate),

		Versions: ToPolicyVersions(p.PolicyVersionList),
		PermissionsBoundaryUsageCount: aws.ToInt32(
			p.PermissionsBoundaryUsageCount),
	}
}

func (r *ManagedPolicyDetailRecorder) Do(ctx context.Context,
	policyDetail types.ManagedPolicyDetail) error {
	return r.Orm.ManagedPolicyDetailModel.Insert(ctx,
		ToManagedPolicyDetail(policyDetail))
}
