package model

import "context"

type ORM struct {
	PolicyModel  PolicyModel
	UserModel    UserModel
	RoleModel    RoleModel
	AccountModel AccountModel
	BucketModel  BucketModel
}

func (o *ORM) Dump(ctx context.Context) {
	o.PolicyModel.Dump(ctx)
	o.UserModel.Dump(ctx)
	o.RoleModel.Dump(ctx)
	o.AccountModel.Dump(ctx)
	o.BucketModel.Dump(ctx)
}
