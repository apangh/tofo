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

type TrailPrintCB struct{}

var _ model.TrailCB = (*TrailPrintCB)(nil)

func (t *TrailPrintCB) Do(ctx context.Context, trail *model.Trail) error {
	fmt.Printf("Trail[%s]:%+v\n", trail.Name, trail)
	return nil
}

func (m *TrailModelMem) Dump(ctx context.Context) error {
	return m.Scan(ctx, &TrailPrintCB{})
}

func (m *TrailModelMem) Scan(ctx context.Context, cb model.TrailCB) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, v := range m.name2Trail {
		if e := cb.Do(ctx, v); e != nil {
			return e
		}
	}
	return nil
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
