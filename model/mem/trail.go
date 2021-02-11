package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type TrailModelMem struct {
	name2Trail map[string]*model.Trail
	mutex      sync.Mutex
}

var _ model.TrailModel = (*TrailModelMem)(nil)

func NewTrailModel() model.TrailModel {
	return &TrailModelMem{
		name2Trail: make(map[string]*model.Trail),
	}
}

func (m *TrailModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.name2Trail {
		fmt.Printf("Trail[%s]:%+v\n", k, v)
	}
}

func (m *TrailModelMem) Lookup(ctx context.Context, name string) (*model.Trail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	t, ok := m.name2Trail[name]
	if ok {
		return t, nil
	}
	return nil, fmt.Errorf("%w: Failed to faind trail '%s'",
		model.TrailNotFound, name)
}

func (m *TrailModelMem) Insert(ctx context.Context, t *model.Trail) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.name2Trail[t.Name]
	if ok {
		return fmt.Errorf("%w: Trail '%s' already exists",
			model.TrailAlreadyExists, t.Name)
	}
	m.name2Trail[t.Name] = t
	return nil
}
