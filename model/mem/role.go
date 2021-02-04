package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type RoleModelMem struct {
	id2Role map[string]*model.Role
	mutex   sync.Mutex
}

var _ model.RoleModel = (*RoleModelMem)(nil)

func NewRoleModel() model.RoleModel {
	return &RoleModelMem{
		id2Role: make(map[string]*model.Role),
	}
}

func (m *RoleModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Role {
		fmt.Printf("Role[%s]:%+v\n", k, v)
	}
}

func (m *RoleModelMem) Lookup(ctx context.Context, id string) (
	*model.Role, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	r, ok := m.id2Role[id]
	if ok {
		return r, nil
	}
	return nil, fmt.Errorf("%w: Failed to find role '%s'",
		model.RoleNotFound, id)
}

func (m *RoleModelMem) Insert(ctx context.Context, r *model.Role) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Role[r.Id]
	if ok {
		return fmt.Errorf("%w: Role %s already exists",
			model.RoleAlreadyExists, r.Id)
	}
	m.id2Role[r.Id] = r
	return nil
}
