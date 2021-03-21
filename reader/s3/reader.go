package s3reader

import (
	"context"
	"io"
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
	key    string
}

var _ reader.Reader = (*s3Reader)(nil)

// NewReader creates a s3 reader
func NewReader(client *s3.Client, bucket, key string) reader.Reader {
	return &s3Reader{
		client: client,
		bucket: bucket,
		key:    key,
	}
}

// Do performs read operation on a s3 object
func (r *s3Reader) Do(ctx context.Context, i reader.Input) (reader.Output, error) {
	start := time.Now()
	rnge := aws.String(
		"bytes=" + strconv.FormatUint(i.GetOffset(), 10) + "-" +
			strconv.FormatUint(i.GetOffset()+i.GetBufs().GetSize()-1, 10))
	params := s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(r.key),
		Range:  rnge,
	}
	o, e := r.client.GetObject(ctx, &params)
	elapsed := time.Since(start)
	if e != nil {
		awslogger.LogAPIError(ctx, e, "Get object %s/%s", r.bucket, r.key)
		logger.Errorf(ctx, "Get object %s/%s failed - elapsed: %+v", r.bucket,
			r.key, elapsed)
		e1 := backtrace.NewErr(e)
		i.RegisterError(ctx, e1)
		return nil, e1
	}
	readSize, e := i.GetBufs().ReadAll(ctx, o.Body)
	if e != nil && e != io.EOF {
		i.RegisterError(ctx, e)
		return nil, e
	}
	oo, e := i.MakeOutput(ctx, readSize, elapsed)
	if e != nil {
		i.RegisterError(ctx, e)
		return nil, e
	}
	if om, ok := oo.(reader.OutputWithMetadata); ok {
		om.SetMetadata(o.Metadata)
	}
	return oo, nil
}

// GetSize return the size of the s3 object
func (r *s3Reader) GetSize(ctx context.Context) (uint64, error) {
	params := s3.HeadObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(r.key),
	}
	o, e := r.client.HeadObject(ctx, &params)
	if e != nil {
		return 0, e
	}
	return uint64(o.ContentLength), nil
}
