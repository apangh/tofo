package mem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/apangh/tofo/model"
)

type BucketModelMem struct {
	buckets map[string]*model.Bucket
	mutex   sync.Mutex
}

var _ model.BucketModel = (*BucketModelMem)(nil)

func (m *BucketModelMem) Insert(ctx context.Context, name string,
	creationDate time.Time, owner *model.User) (*model.Bucket, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.buckets[name]
	if ok {
		return nil, fmt.Errorf("%w: Bucket %s already exists",
			model.BucketAlreadyExist, name)
	}
	b := &model.Bucket{
		Name:         name,
		CreationDate: creationDate,
		Owner:        owner,
	}
	m.buckets[name] = b
	return b, nil
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
