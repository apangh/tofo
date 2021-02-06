package mem

import "github.com/apangh/tofo/model"

func NewORM() (*model.ORM, error) {
	return &model.ORM{
		ManagedPolicyModel:       NewManagedPolicyModel(),
		ManagedPolicyDetailModel: NewManagedPolicyDetailModel(),
		RoleModel:                NewRoleModel(),
		GroupModel:               NewGroupModel(),
		UserModel:                NewUserModel(),
		AccountModel:             NewAccountModel(),
		BucketModel:              NewBucketModel(),
	}, nil
}
