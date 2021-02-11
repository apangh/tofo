package recorder

import (
	"github.com/apangh/tofo/model"
	commonRecorder "github.com/apangh/tofo/recorder"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

func ToTrail(t *types.Trail) (*model.Trail, error) {
	if t == nil {
		return nil, nil
	}

	arn, e := commonRecorder.ToArn(t.TrailARN)
	if e != nil {
		return nil, e
	}

	var cwLogsGroupArn *model.ARN
	if t.CloudWatchLogsLogGroupArn != nil {
		cwLogsGroupArn, e = commonRecorder.ToArn(t.CloudWatchLogsLogGroupArn)
		if e != nil {
			return nil, e
		}
	}
	var cwLogsRoleArn *model.ARN
	if t.CloudWatchLogsRoleArn != nil {
		cwLogsRoleArn, e = commonRecorder.ToArn(t.CloudWatchLogsRoleArn)
		if e != nil {
			return nil, e
		}
	}
	var kmsKeyIdArn *model.ARN
	if t.KmsKeyId != nil {
		kmsKeyIdArn, e = commonRecorder.ToArn(t.KmsKeyId)
		if e != nil {
			return nil, e
		}
	}
	var snsTopicArn *model.ARN
	if t.SnsTopicARN != nil {
		snsTopicArn, e = commonRecorder.ToArn(t.SnsTopicARN)
		if e != nil {
			return nil, e
		}
	}

	return &model.Trail{
		Arn:                        arn,
		Name:                       aws.ToString(t.Name),
		HomeRegion:                 aws.ToString(t.HomeRegion),
		S3BucketName:               aws.ToString(t.S3BucketName),
		S3KeyPrefix:                aws.ToString(t.S3KeyPrefix),
		CWLogsGroupArn:             cwLogsGroupArn,
		CWLogsRoleArn:              cwLogsRoleArn,
		KmsKeyIdArn:                kmsKeyIdArn,
		SnsTopicArn:                snsTopicArn,
		HasCustomEventSelectors:    aws.ToBool(t.HasCustomEventSelectors),
		HasInsightSelectors:        aws.ToBool(t.HasInsightSelectors),
		IncludeGlobalServiceEvents: aws.ToBool(t.IncludeGlobalServiceEvents),
		IsMultiRegionTrail:         aws.ToBool(t.IsMultiRegionTrail),
		IsOrganizationTrail:        aws.ToBool(t.IsOrganizationTrail),
		LogFileVallidationEnabled:  aws.ToBool(t.LogFileValidationEnabled),
	}, nil
}
