package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type ManagedPolicyModelMem struct {
	id2Policy  map[string]*model.ManagedPolicy
	arn2Policy map[string]*model.ManagedPolicy
	mutex      sync.Mutex
}

var _ model.ManagedPolicyModel = (*ManagedPolicyModelMem)(nil)

func NewManagedPolicyModel() model.ManagedPolicyModel {
	return &ManagedPolicyModelMem{
		id2Policy:  make(map[string]*model.ManagedPolicy),
		arn2Policy: make(map[string]*model.ManagedPolicy),
	}
}

func (m *ManagedPolicyModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Policy {
		fmt.Printf("ManagedPolicy[%s]:%+v\n", k, v)
	}
}

func (m *ManagedPolicyModelMem) Lookup(ctx context.Context, id string) (
	*model.ManagedPolicy, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.id2Policy[id]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find managed policy by id '%s'",
		model.ManagedPolicyNotFound, id)
}

func (m *ManagedPolicyModelMem) LookupByArn(ctx context.Context, arn string) (
	*model.ManagedPolicy, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.arn2Policy[arn]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find managed policy by arn '%s'",
		model.ManagedPolicyNotFound, arn)
}

func (m *ManagedPolicyModelMem) Insert(ctx context.Context, p *model.ManagedPolicy) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Policy[p.Id]
	if ok {
		return fmt.Errorf("%w: Managed Policy %s already exists",
			model.ManagedPolicyAlreadyExists, p.Id)
	}
	m.id2Policy[p.Id] = p
	m.arn2Policy[p.Arn] = p
	return nil
}
