package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type RoleDetailModelMem struct {
	id2Role map[string]*model.RoleDetail
	mutex   sync.Mutex
}

var _ model.RoleDetailModel = (*RoleDetailModelMem)(nil)

func NewRoleDetailModel() model.RoleDetailModel {
	return &RoleDetailModelMem{
		id2Role: make(map[string]*model.RoleDetail),
	}
}

func (m *RoleDetailModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Role {
		fmt.Printf("RoleDetail[%s]:%+v\n", k, v)
	}
}

func (m *RoleDetailModelMem) Lookup(ctx context.Context, id string) (
	*model.RoleDetail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	r, ok := m.id2Role[id]
	if ok {
		return r, nil
	}
	return nil, fmt.Errorf("%w: Failed to find role detail '%s'",
		model.RoleDetailNotFound, id)
}

func (m *RoleDetailModelMem) Insert(ctx context.Context, r *model.RoleDetail) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Role[r.Id]
	if ok {
		return fmt.Errorf("%w: Role detail %s already exists",
			model.RoleDetailAlreadyExists, r.Id)
	}
	m.id2Role[r.Id] = r
	return nil
}
