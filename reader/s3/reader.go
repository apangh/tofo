package s3reader

import (
	"context"
	"strconv"
	"time"

	"github.com/apangh/salt/debug/backtrace"
	"github.com/apangh/salt/logger"
	"github.com/apangh/salt/reader"
	awslogger "github.com/apangh/tofo/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// s3 reader
type s3Reader struct {
	client *s3.Client
	bucket string
}

var _ reader.Reader = (*s3Reader)(nil)

// NewReader creates a s3 reader
func NewReader(client *s3.Client, bucket string) reader.Reader {
	return &s3Reader{
		client: client,
		bucket: bucket,
	}
}

// Do performs read operation on a s3 object
func (r *s3Reader) Do(ctx context.Context, i reader.Input) (reader.Output, error) {
	i2 := i.(reader.InputWithKey)
	start := time.Now()
	rnge := aws.String(
		"bytes=" + strconv.FormatUint(i.GetOffset(), 10) + "-" +
			strconv.FormatUint(i.GetOffset()+i.GetBufs().GetSize()-1, 10))
	params := s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(i2.GetKey()),
		Range:  rnge,
	}
	o, e := r.client.GetObject(ctx, &params)
	elapsed := time.Since(start)
	if e != nil {
		awslogger.LogAPIError(ctx, e, "Get object %s/%s", r.bucket, i2.GetKey())
		logger.Errorf(ctx, "Get object %s/%s failed - elapsed: %+v", r.bucket,
			i2.GetKey(), elapsed)
		return nil, backtrace.NewErr(e)
	}
	readSize, e := i.GetBufs().ReadAll(ctx, o.Body)
	oo := i.MakeOutput(readSize, elapsed)
	if om, ok := oo.(reader.OutputWithMetadata); ok {
		om.SetMetadata(o.Metadata)
	}
	return oo, e
}
