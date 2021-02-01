package dynamodbutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/glog"
)

var _ BackupSummaryCB = (*LogBackupSummary)(nil)

type LogBackupSummary struct {
	i int
}

func (s *LogBackupSummary) Do(ctx context.Context, summary types.BackupSummary) error {
	glog.Infof("Backup[%d]: %s %s %v %v %d %v %v", s.i,
		aws.ToString(summary.BackupArn),
		aws.ToString(summary.BackupName),
		summary.BackupCreationDateTime,
		summary.BackupExpiryDateTime,
		aws.ToInt64(summary.BackupSizeBytes),
		summary.BackupStatus,
		summary.BackupType)
	s.i++
	return nil
}
