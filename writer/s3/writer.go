// This file implements writer for S3 objects
package s3writer

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/apangh/salt/logger"
	"github.com/apangh/salt/memory/mpool"
	"github.com/apangh/salt/writer"
	awslogger "github.com/apangh/tofo/logger"
)

// s3 writer
type s3Writer struct {
	client *s3.Client
	bucket string
}

var _ writer.Writer = (*s3Writer)(nil)

// NewWriter creates a s3 writer
func NewWriter(client *s3.Client, bucket string) writer.Writer {
	return &s3Writer{
		client: client,
		bucket: bucket,
	}
}

// do is an internal implementation of Do
func (w *s3Writer) do(ctx context.Context, i writer.Input) (writer.Output, error) {
	start := time.Now()
	s3tags := url.Values{}
	for k, v := range i.GetTags() {
		s3tags.Add(k, v)
	}
	key := i.GetKey()
	params := s3.PutObjectInput{
		Bucket:   aws.String(w.bucket),
		Key:      aws.String(key),
		Body:     mpool.NewBufsReadSeeker(i.GetBufs()),
		Metadata: i.GetMetadata(),
		Tagging:  aws.String(s3tags.Encode()),
	}
	size := i.GetBufs().GetSize()
	o, e := w.client.PutObject(ctx, &params)
	elapsed := time.Since(start)
	if e != nil {
		awslogger.LogAPIError(ctx, e, "Put object %s/%s", w.bucket, key)
		logger.Errorf(ctx, "Put object %s/%s failed - elapsed: %+v", w.bucket, key, elapsed)
		return nil, e
	}
	output, e := i.MakeOutput(ctx, elapsed)
	if e != nil {
		return nil, e
	}
	requestId, _ := middleware.GetRequestIDMetadata(o.ResultMetadata)
	logger.Infof(ctx, "Put object %s/%s, requestID: %s, elapsed: %+v, throughput: %.2f B/s",
		w.bucket, key, requestId, elapsed, float64(size)/elapsed.Seconds())
	return output, nil
}

// Do creates a S3 object with content from bufs
func (w *s3Writer) Do(ctx context.Context, i writer.Input) (writer.Output, error) {
	o, e := w.do(ctx, i)
	if e != nil {
		i.RegisterError(ctx, e)
		return nil, e
	}
	return o, nil
}
