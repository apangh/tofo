package logger

import (
	"context"
	"errors"
	"fmt"

	"github.com/apangh/salt/logger"
	"github.com/aws/smithy-go"
)

func LogAPIError(ctx context.Context, e error, format string,
	args ...interface{}) {
	prefix := fmt.Sprintf(format, args...)
	var ae smithy.APIError
	if errors.As(e, &ae) {
		logger.Errorf(ctx, "%s failed - code: %s, message: %s, fault: %s", prefix,
			ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
	}
	var re interface {
		ServiceHostID() string
		ServiceRequestID() string
	}
	if errors.As(e, &re) {
		logger.Errorf(ctx, "%s failed - requestID: %s, hostID: %s", prefix,
			re.ServiceRequestID(), re.ServiceHostID())
	}
}
