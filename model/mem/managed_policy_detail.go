package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type ManagedPolicyDetailModelMem struct {
	id2Policy  map[string]*model.ManagedPolicyDetail
	arn2Policy map[string]*model.ManagedPolicyDetail
	mutex      sync.Mutex
}

var _ model.ManagedPolicyDetailModel = (*ManagedPolicyDetailModelMem)(nil)

func NewManagedPolicyDetailModel() model.ManagedPolicyDetailModel {
	return &ManagedPolicyDetailModelMem{
		id2Policy:  make(map[string]*model.ManagedPolicyDetail),
		arn2Policy: make(map[string]*model.ManagedPolicyDetail),
	}
}

func (m *ManagedPolicyDetailModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.id2Policy {
		fmt.Printf("ManagedPolicyDetail[%s]:%+v\n", k, v)
	}
}

func (m *ManagedPolicyDetailModelMem) Lookup(ctx context.Context, id string) (
	*model.ManagedPolicyDetail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.id2Policy[id]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find managed policy detail by id '%s'",
		model.ManagedPolicyDetailNotFound, id)
}

func (m *ManagedPolicyDetailModelMem) LookupByArn(ctx context.Context, arn string) (
	*model.ManagedPolicyDetail, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	p, ok := m.arn2Policy[arn]
	if ok {
		return p, nil
	}
	return nil, fmt.Errorf("%w: Failed to find managed policy detail by arn '%s'",
		model.ManagedPolicyDetailNotFound, arn)
}

func (m *ManagedPolicyDetailModelMem) Insert(ctx context.Context,
	p *model.ManagedPolicyDetail) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.id2Policy[p.Id]
	if ok {
		return fmt.Errorf("%w: Managed Policy detail %s already exists",
			model.ManagedPolicyDetailAlreadyExists, p.Id)
	}
	m.id2Policy[p.Id] = p
	m.arn2Policy[p.Arn.String()] = p
	return nil
}
