package mem

import "github.com/apangh/tofo/model"

func NewORM() (*model.ORM, error) {
	return &model.ORM{
		BucketModel:  NewBucketModel(),
		UserModel:    NewUserModel(),
		AccountModel: NewAccountModel(),
	}, nil
}
