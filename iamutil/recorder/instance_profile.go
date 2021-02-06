package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToInstanceProfile(i types.InstanceProfile) *model.InstanceProfile {
	return &model.InstanceProfile{
		Arn:        aws.ToString(i.Arn),
		CreateDate: aws.ToTime(i.CreateDate),
		Id:         aws.ToString(i.InstanceProfileId),
		Name:       aws.ToString(i.InstanceProfileName),
		Path:       aws.ToString(i.Path),
		Roles:      ToRoles(i.Roles),
	}
}

func ToInstanceProfiles(i []types.InstanceProfile) []*model.InstanceProfile {
	res := make([]*model.InstanceProfile, 0, len(i))
	for _, a := range i {
		res = append(res, ToInstanceProfile(a))
	}
	return res
}
