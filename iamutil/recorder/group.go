package recorder

import (
	"github.com/apangh/tofo/model"
	commonRecorder "github.com/apangh/tofo/recorder"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToGroup(group types.Group) (*model.Group, error) {
	arn, e := commonRecorder.ToArn(group.Arn)
	if e != nil {
		return nil, e
	}
	return &model.Group{
		Id:         aws.ToString(group.GroupId),
		Name:       aws.ToString(group.GroupName),
		Path:       aws.ToString(group.Path),
		Arn:        arn,
		CreateDate: aws.ToTime(group.CreateDate),
	}, nil
}
