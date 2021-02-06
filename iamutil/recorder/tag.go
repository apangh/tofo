package recorder

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ToTags(tags []types.Tag) map[string]string {
	res := make(map[string]string)
	for _, t := range tags {
		res[*t.Key] = *t.Value
	}
	return res
}
