package model

import "context"

type ORM struct {
	ManagedPolicyModel       ManagedPolicyModel
	ManagedPolicyDetailModel ManagedPolicyDetailModel
	GroupModel               GroupModel
	UserModel                UserModel
	RoleModel                RoleModel
	AccountModel             AccountModel
	BucketModel              BucketModel
}

func (o *ORM) Dump(ctx context.Context) {
	o.ManagedPolicyModel.Dump(ctx)
	o.ManagedPolicyDetailModel.Dump(ctx)
	o.GroupModel.Dump(ctx)
	o.UserModel.Dump(ctx)
	o.RoleModel.Dump(ctx)
	o.AccountModel.Dump(ctx)
	o.BucketModel.Dump(ctx)
}
