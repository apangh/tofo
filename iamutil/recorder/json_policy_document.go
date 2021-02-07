package recorder

import (
	"net/url"

	"github.com/apangh/tofo/model"
)

func ToJsonPolicyDocument(s string) (model.JsonPolicyDocument, error) {
	res, e := url.QueryUnescape(s)
	if e != nil {
		return "", e
	}
	return (model.JsonPolicyDocument)(res), nil
}
