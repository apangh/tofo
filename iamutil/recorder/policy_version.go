package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func toPolicyVersion(pv types.PolicyVersion) *model.PolicyVersion {
	return &model.PolicyVersion{
		IsDefaultVersion: pv.IsDefaultVersion,
		CreateDate:       aws.ToTime(pv.CreateDate),
		Document:         aws.ToString(pv.Document),
		VersionId:        aws.ToString(pv.VersionId),
	}
}

func ToPolicyVersions(p []types.PolicyVersion) map[string]*model.PolicyVersion {
	res := make(map[string]*model.PolicyVersion)
	for _, v1 := range p {
		pv := toPolicyVersion(v1)
		res[pv.VersionId] = pv
	}
	return res
}
