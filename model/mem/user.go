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

func (m *UserModelMem) Insert(ctx context.Context, id, displayName string) (
	*model.User, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2User[id]
	if ok {
		return nil, fmt.Errorf("%w: User %s/%s already exists",
			model.UserAlreadyExists, id, displayName)
	}
	u := &model.User{
		Id:          id,
		DisplayName: displayName,
	}
	m.id2User[id] = u
	return u, nil
}
