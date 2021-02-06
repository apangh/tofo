package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToInlinePolicyDetail(a types.PolicyDetail) *model.InlinePolicyDetail {
	return &model.InlinePolicyDetail{
		Name:     aws.ToString(a.PolicyName),
		Document: aws.ToString(a.PolicyDocument),
	}
}

func ToInlinePolicyDetails(a []types.PolicyDetail) []*model.InlinePolicyDetail {
	res := make([]*model.InlinePolicyDetail, 0, len(a))
	for _, p := range a {
		res = append(res, ToInlinePolicyDetail(p))
	}
	return res
}
