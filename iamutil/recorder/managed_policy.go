package recorder

import (
	"context"

	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type ManagedPolicyRecorder struct {
	Orm *model.ORM
}

func toManagedPolicy(p types.Policy) (*model.ManagedPolicy, error) {
	arn, e := ToArn(p.Arn)
	if e != nil {
		return nil, e
	}
	return &model.ManagedPolicy{
		Id:                            aws.ToString(p.PolicyId),
		Name:                          aws.ToString(p.PolicyName),
		Path:                          aws.ToString(p.Path),
		Arn:                           arn,
		AttachmentCount:               aws.ToInt32(p.AttachmentCount),
		PermissionsBoundaryUsageCount: aws.ToInt32(p.PermissionsBoundaryUsageCount),
		DefaultVersionId:              aws.ToString(p.DefaultVersionId),
		Description:                   aws.ToString(p.Description),
		IsAttachable:                  p.IsAttachable,
		CreateDate:                    aws.ToTime(p.CreateDate),
		UpdateDate:                    aws.ToTime(p.UpdateDate),
	}, nil
}

func (r *ManagedPolicyRecorder) Do(ctx context.Context, policy types.Policy) error {
	p, e := toManagedPolicy(policy)
	if e != nil {
		return e
	}
	return r.Orm.ManagedPolicyModel.Insert(ctx, p)
}
