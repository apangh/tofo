package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type PolicyModelMem struct {
	id2Policy  map[string]*model.Policy
	arn2Policy map[string]*model.Policy
	mutex      sync.Mutex
}

var _ model.PolicyModel = (*PolicyModelMem)(nil)

func NewPolicyModel() model.PolicyModel {
	return &PolicyModelMem{
		id2Policy:  make(map[string]*model.Policy),
		arn2Policy: make(map[string]*model.Policy),
	}
}

func (m *PolicyModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Policy {
		fmt.Printf("Policy[%s]:%+v\n", k, v)
	}
}

func (m *PolicyModelMem) Lookup(ctx context.Context, id string) (
	*model.Policy, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.id2Policy[id]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find policy by id '%s'",
		model.PolicyNotFound, id)
}

func (m *PolicyModelMem) LookupByArn(ctx context.Context, arn string) (
	*model.Policy, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.arn2Policy[arn]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find policy by arn '%s'",
		model.PolicyNotFound, arn)
}

func (m *PolicyModelMem) Insert(ctx context.Context, p *model.Policy) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Policy[p.Id]
	if ok {
		return fmt.Errorf("%w: Policy %s already exists",
			model.PolicyAlreadyExists, p.Id)
	}
	m.id2Policy[p.Id] = p
	m.arn2Policy[p.Arn] = p
	return nil
}
