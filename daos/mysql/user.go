package mysql

type IUserModel interface {
	//判断用户是否存在
	CheckUserExist(username string) bool
}

type UserModel struct {
}

func NewUserModel() IUserModel {
	return &UserModel{}
}

func (this *UserModel) CheckUserExist(username string) bool {
	return false
}
