package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type GroupModelMem struct {
	id2Group map[string]*model.Group
	mutex    sync.Mutex
}

var _ model.GroupModel = (*GroupModelMem)(nil)

func NewGroupModel() model.GroupModel {
	return &GroupModelMem{
		id2Group: make(map[string]*model.Group),
	}
}

func (m *GroupModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Group {
		fmt.Printf("Group[%s]:%+v\n", k, v)
	}
}

func (m *GroupModelMem) Lookup(ctx context.Context, id string) (
	*model.Group, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	g, ok := m.id2Group[id]
	if ok {
		return g, nil
	}
	return nil, fmt.Errorf("%w: Failed to find group '%s'",
		model.GroupNotFound, id)
}

func (m *GroupModelMem) Insert(ctx context.Context, g *model.Group) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Group[g.Id]
	if ok {
		return fmt.Errorf("%w: Group %s already exists",
			model.GroupAlreadyExists, g.Id)
	}
	m.id2Group[g.Id] = g
	return nil
}
