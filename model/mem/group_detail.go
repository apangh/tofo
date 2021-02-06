package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type GroupDetailModelMem struct {
	id2Group map[string]*model.GroupDetail
	mutex    sync.Mutex
}

var _ model.GroupDetailModel = (*GroupDetailModelMem)(nil)

func NewGroupDetailModel() model.GroupDetailModel {
	return &GroupDetailModelMem{
		id2Group: make(map[string]*model.GroupDetail),
	}
}

func (m *GroupDetailModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Group {
		fmt.Printf("GroupDetail[%s]:%+v\n", k, v)
	}
}

func (m *GroupDetailModelMem) Lookup(ctx context.Context, id string) (
	*model.GroupDetail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	r, ok := m.id2Group[id]
	if ok {
		return r, nil
	}
	return nil, fmt.Errorf("%w: Failed to find group detail '%s'",
		model.GroupDetailNotFound, id)
}

func (m *GroupDetailModelMem) Insert(ctx context.Context, r *model.GroupDetail) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Group[r.Id]
	if ok {
		return fmt.Errorf("%w: Group detail %s already exists",
			model.GroupDetailAlreadyExists, r.Id)
	}
	m.id2Group[r.Id] = r
	return nil
}
