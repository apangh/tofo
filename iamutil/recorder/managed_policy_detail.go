package recorder

import (
	"context"

	"github.com/apangh/tofo/model"
	commonRecorder "github.com/apangh/tofo/recorder"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type ManagedPolicyDetailRecorder struct {
	Orm *model.ORM
}

func ToManagedPolicyDetail(p types.ManagedPolicyDetail) (
	*model.ManagedPolicyDetail, error) {
	versions, e := ToPolicyVersions(p.PolicyVersionList)
	if e != nil {
		return nil, e
	}
	arn, e := commonRecorder.ToArn(p.Arn)
	if e != nil {
		return nil, e
	}
	return &model.ManagedPolicyDetail{
		Id:               aws.ToString(p.PolicyId),
		Name:             aws.ToString(p.PolicyName),
		Path:             aws.ToString(p.Path),
		Arn:              arn,
		AttachmentCount:  aws.ToInt32(p.AttachmentCount),
		DefaultVersionId: aws.ToString(p.DefaultVersionId),
		Description:      aws.ToString(p.Description),
		IsAttachable:     p.IsAttachable,
		CreateDate:       aws.ToTime(p.CreateDate),
		UpdateDate:       aws.ToTime(p.UpdateDate),
		Versions:         versions,
		PermissionsBoundaryUsageCount: aws.ToInt32(
			p.PermissionsBoundaryUsageCount),
	}, nil
}

func (r *ManagedPolicyDetailRecorder) Do(ctx context.Context,
	policyDetail types.ManagedPolicyDetail) error {
	d, e := ToManagedPolicyDetail(policyDetail)
	if e != nil {
		return e
	}
	return r.Orm.ManagedPolicyDetailModel.Insert(ctx, d)
}
