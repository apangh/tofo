package recorder

import (
	"encoding/json"
	"net/url"

	"github.com/apangh/tofo/model"
)

func ToJsonPolicyDocument(s string) (*model.JsonPolicyDocument, error) {
	res, e := url.QueryUnescape(s)
	if e != nil {
		return nil, e
	}
	var pd model.JsonPolicyDocument
	if e := json.Unmarshal([]byte(res), &pd); e != nil {
		return nil, e
	}
	return &pd, nil
}
