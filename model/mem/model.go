package mem

import "github.com/apangh/tofo/model"

func NewORM() (*model.ORM, error) {
	return &model.ORM{
		ManagedPolicyDetailModel: NewManagedPolicyDetailModel(),
		GroupDetailModel:         NewGroupDetailModel(),
		UserDetailModel:          NewUserDetailModel(),
		RoleDetailModel:          NewRoleDetailModel(),
		ManagedPolicyModel:       NewManagedPolicyModel(),
		RoleModel:                NewRoleModel(),
		GroupModel:               NewGroupModel(),
		UserModel:                NewUserModel(),
		AccountModel:             NewAccountModel(),
		BucketModel:              NewBucketModel(),
		TrailModel:               NewTrailModel(),
	}, nil
}
