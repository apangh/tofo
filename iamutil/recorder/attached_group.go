package recorder

import "github.com/apangh/tofo/model"

func ToAttachedGroup(groupName string) *model.AttachedGroup {
	return &model.AttachedGroup{
		Name: groupName,
	}
}

func ToAttachedGroups(names []string) []*model.AttachedGroup {
	res := make([]*model.AttachedGroup, 0, len(names))
	for _, n := range names {
		res = append(res, ToAttachedGroup(n))
	}
	return res
}
