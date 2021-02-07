package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func toPolicyVersion(pv types.PolicyVersion) (*model.PolicyVersion, error) {
	s, e := ToJsonPolicyDocument(aws.ToString(pv.Document))
	if e != nil {
		return nil, e
	}
	return &model.PolicyVersion{
		IsDefaultVersion: pv.IsDefaultVersion,
		CreateDate:       aws.ToTime(pv.CreateDate),
		Document:         s,
		VersionId:        aws.ToString(pv.VersionId),
	}, nil
}

func ToPolicyVersions(p []types.PolicyVersion) (map[string]*model.PolicyVersion, error) {
	res := make(map[string]*model.PolicyVersion)
	for _, v1 := range p {
		pv, e := toPolicyVersion(v1)
		if e != nil {
			return nil, e
		}
		res[pv.VersionId] = pv
	}
	return res, nil
}
