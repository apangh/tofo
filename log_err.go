package tofo

import (
	"errors"

	"github.com/aws/smithy-go"
	"github.com/golang/glog"
)

func LogErr(prefix string, e error) {
	var ae smithy.APIError
	if errors.As(e, &ae) {
		glog.Errorf("%s code: %s, message: %s, fault: %s",
			prefix, ae.ErrorCode(), ae.ErrorMessage(),
			ae.ErrorFault().String())
	}
	var re interface {
		ServiceHostID() string
		ServiceRequestID() string
	}
	if errors.As(e, &re) {
		glog.Errorf("%s requestID: %s, hostID %s",
			prefix, re.ServiceRequestID(),
			re.ServiceHostID())
	}
}
