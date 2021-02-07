package model

import "github.com/aws/aws-sdk-go-v2/aws/arn"

type ARN struct {
	arn.ARN
}

func (r *ARN) String() string {
	return r.ARN.String()
}
