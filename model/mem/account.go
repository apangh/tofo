package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type AccountModelMem struct {
	id2Account map[string]*model.Account
	mutex      sync.Mutex
}

var _ model.AccountModel = (*AccountModelMem)(nil)

func NewAccountModel() model.AccountModel {
	return &AccountModelMem{
		id2Account: make(map[string]*model.Account),
	}
}

func (m *AccountModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Account {
		fmt.Printf("Account[%s]:%+v\n", k, v)
	}
}

func (m *AccountModelMem) Lookup(ctx context.Context, id string) (
	*model.Account, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	a, ok := m.id2Account[id]
	if ok {
		return a, nil
	}
	return nil, fmt.Errorf("%w: Failed to find account '%s'",
		model.AccountNotFound, id)
}

func (m *AccountModelMem) Insert(ctx context.Context, a *model.Account) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Account[a.Id]
	if ok {
		return fmt.Errorf("%w: Account '%s' already exists",
			model.AccountAlreadyExists, a.Id)
	}
	m.id2Account[a.Id] = a
	return nil
}
