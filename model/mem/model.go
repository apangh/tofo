package mem

import "github.com/apangh/tofo/model"

func NewORM() (*model.ORM, error) {
	return &model.ORM{
		BucketModel: &BucketModelMem{},
		UserModel:   &UserModelMem{},
	}, nil
}
