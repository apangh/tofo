package recorder

import (
	"github.com/apangh/tofo/model"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

func ToArn(s *string) (*model.ARN, error) {
	arn, e := arn.Parse(*s)
	if e != nil {
		return nil, e
	}
	return &model.ARN{
		ARN: arn,
	}, nil
}
