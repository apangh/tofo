package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToGroup(group types.Group) *model.Group {
	return &model.Group{
		Id:         aws.ToString(group.GroupId),
		Name:       aws.ToString(group.GroupName),
		Path:       aws.ToString(group.Path),
		Arn:        aws.ToString(group.Arn),
		CreateDate: aws.ToTime(group.CreateDate),
	}
}
