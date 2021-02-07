package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToAttachedPolicy(a *types.AttachedPolicy) (*model.AttachedPolicy, error) {
	if a == nil {
		return nil, nil
	}
	arn, e := ToArn(a.PolicyArn)
	if e != nil {
		return nil, e
	}
	return &model.AttachedPolicy{
		Arn:  arn,
		Name: a.PolicyName,
	}, nil
}

func ToAttachedPolicies(a []types.AttachedPolicy) ([]*model.AttachedPolicy, error) {
	res := make([]*model.AttachedPolicy, 0, len(a))
	for _, p := range a {
		ap, e := ToAttachedPolicy(&p)
		if e != nil {
			return nil, e
		}
		res = append(res, ap)
	}
	return res, nil
}
