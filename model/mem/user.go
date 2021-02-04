package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type UserModelMem struct {
	id2User map[string]*model.User
	mutex   sync.Mutex
}

var _ model.UserModel = (*UserModelMem)(nil)

func NewUserModel() model.UserModel {
	return &UserModelMem{
		id2User: make(map[string]*model.User),
	}
}

func (m *UserModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2User {
		fmt.Printf("User[%s]:%+v\n", k, v)
	}
}

func (m *UserModelMem) Lookup(ctx context.Context, id string) (
	*model.User, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	u, ok := m.id2User[id]
	if ok {
		return u, nil
	}
	return nil, fmt.Errorf("%w: Failed to find user '%s'",
		model.UserNotFound, id)
}

func (m *UserModelMem) Insert(ctx context.Context, u *model.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2User[u.Id]
	if ok {
		return fmt.Errorf("%w: User %s already exists",
			model.UserAlreadyExists, u.Id)
	}
	m.id2User[u.Id] = u
	return nil
}
