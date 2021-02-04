package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/apangh/tofo/model"
)

type BucketModelMem struct {
	buckets map[string]*model.Bucket
	mutex   sync.Mutex
}

var _ model.BucketModel = (*BucketModelMem)(nil)

func NewBucketModel() model.BucketModel {
	return &BucketModelMem{
		buckets: make(map[string]*model.Bucket),
	}
}

func (m *BucketModelMem) Dump(ctx context.Context) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for k, v := range m.buckets {
		fmt.Printf("Bucket[%s]:%+v\n", k, v)
	}
}

func (m *BucketModelMem) Insert(ctx context.Context, b *model.Bucket) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.buckets[b.Name]
	if ok {
		return fmt.Errorf("%w: Bucket %s already exists",
			model.BucketAlreadyExist, b.Name)
	}
	m.buckets[b.Name] = b
	return nil
}

func (m *BucketModelMem) Lookup(ctx context.Context, name string) (
	*model.Bucket, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	b, ok := m.buckets[name]
	if !ok {
		return nil, fmt.Errorf("%w: Bucket %s not found",
			model.BucketNotFound, name)
	}
	return b, nil
}
