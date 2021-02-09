package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToInlinePolicyDetail(a types.PolicyDetail) (*model.InlinePolicyDetail, error) {
	s, e := ToIamPolicyDocument(aws.ToString(a.PolicyDocument))
	if e != nil {
		return nil, e
	}
	return &model.InlinePolicyDetail{
		Name:     aws.ToString(a.PolicyName),
		Document: s,
	}, nil
}

func ToInlinePolicyDetails(a []types.PolicyDetail) ([]*model.InlinePolicyDetail, error) {
	res := make([]*model.InlinePolicyDetail, 0, len(a))
	for _, p := range a {
		p1, e := ToInlinePolicyDetail(p)
		if e != nil {
			return nil, e
		}
		res = append(res, p1)
	}
	return res, nil
}
