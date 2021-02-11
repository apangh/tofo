package model

import "context"

type ORM struct {
	ManagedPolicyDetailModel ManagedPolicyDetailModel
	GroupDetailModel         GroupDetailModel
	UserDetailModel          UserDetailModel
	RoleDetailModel          RoleDetailModel
	ManagedPolicyModel       ManagedPolicyModel
	GroupModel               GroupModel
	UserModel                UserModel
	RoleModel                RoleModel
	AccountModel             AccountModel
	BucketModel              BucketModel
	TrailModel               TrailModel
}

func (o *ORM) Dump(ctx context.Context) {
	o.ManagedPolicyDetailModel.Dump(ctx)
	o.GroupDetailModel.Dump(ctx)
	o.RoleDetailModel.Dump(ctx)
	o.UserDetailModel.Dump(ctx)
	o.ManagedPolicyModel.Dump(ctx)
	o.GroupModel.Dump(ctx)
	o.UserModel.Dump(ctx)
	o.RoleModel.Dump(ctx)
	o.AccountModel.Dump(ctx)
	o.BucketModel.Dump(ctx)
	o.TrailModel.Dump(ctx)
}
