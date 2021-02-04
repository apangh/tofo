package model

import "context"

type ORM struct {
	AccountModel AccountModel
	UserModel    UserModel
	BucketModel  BucketModel
}

func (o *ORM) Dump(ctx context.Context) {
	o.AccountModel.Dump(ctx)
	o.UserModel.Dump(ctx)
	o.BucketModel.Dump(ctx)
}
