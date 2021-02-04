package mem

import "github.com/apangh/tofo/model"

func NewORM() (*model.ORM, error) {
	return &model.ORM{
		PolicyModel:  NewPolicyModel(),
		RoleModel:    NewRoleModel(),
		GroupModel:   NewGroupModel(),
		UserModel:    NewUserModel(),
		AccountModel: NewAccountModel(),
		BucketModel:  NewBucketModel(),
	}, nil
}
