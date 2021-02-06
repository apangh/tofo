package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToAttachedPolicy(a *types.AttachedPolicy) *model.AttachedPolicy {
	if a == nil {
		return nil
	}
	return &model.AttachedPolicy{
		Arn:  a.PolicyArn,
		Name: a.PolicyName,
	}
}

func ToAttachedPolicies(a []types.AttachedPolicy) []*model.AttachedPolicy {
	res := make([]*model.AttachedPolicy, 0, len(a))
	for _, p := range a {
		res = append(res, ToAttachedPolicy(&p))
	}
	return res
}
