package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type UserDetailModelMem struct {
	id2User map[string]*model.UserDetail
	mutex   sync.Mutex
}

var _ model.UserDetailModel = (*UserDetailModelMem)(nil)

func NewUserDetailModel() model.UserDetailModel {
	return &UserDetailModelMem{
		id2User: make(map[string]*model.UserDetail),
	}
}

func (m *UserDetailModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2User {
		fmt.Printf("UserDetail[%s]:%+v\n", k, v)
	}
}

func (m *UserDetailModelMem) Lookup(ctx context.Context, id string) (
	*model.UserDetail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	r, ok := m.id2User[id]
	if ok {
		return r, nil
	}
	return nil, fmt.Errorf("%w: Failed to find user detail '%s'",
		model.UserDetailNotFound, id)
}

func (m *UserDetailModelMem) Insert(ctx context.Context, r *model.UserDetail) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2User[r.Id]
	if ok {
		return fmt.Errorf("%w: User detail %s already exists",
			model.UserDetailAlreadyExists, r.Id)
	}
	m.id2User[r.Id] = r
	return nil
}
