package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToInstanceProfile(i types.InstanceProfile) (*model.InstanceProfile, error) {
	arn, e := ToArn(i.Arn)
	if e != nil {
		return nil, e
	}
	roles, e := ToRoles(i.Roles)
	if e != nil {
		return nil, e
	}
	return &model.InstanceProfile{
		Arn:        arn,
		CreateDate: aws.ToTime(i.CreateDate),
		Id:         aws.ToString(i.InstanceProfileId),
		Name:       aws.ToString(i.InstanceProfileName),
		Path:       aws.ToString(i.Path),
		Roles:      roles,
	}, nil
}

func ToInstanceProfiles(i []types.InstanceProfile) ([]*model.InstanceProfile, error) {
	res := make([]*model.InstanceProfile, 0, len(i))
	for _, a := range i {
		ip, e := ToInstanceProfile(a)
		if e != nil {
			return nil, e
		}
		res = append(res, ip)
	}
	return res, nil
}
