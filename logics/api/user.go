package api

import (
	"blog_tpl/daos/mysql"
	"blog_tpl/models"
	"blog_tpl/utils"
	"fmt"
)

type IUserInfo interface {
	//获取用户信息
	GetUserInfo(id int) models.User

	//注册用户
	Register(id int) bool
}

type UserInfo struct {
	userModel mysql.IUserModel
}

func NewUserInfo() IUserInfo {
	return &UserInfo{
		userModel: mysql.NewUserModel(),
	}
}

//获取用户信息
func (this *UserInfo) GetUserInfo(id int) models.User {
	return models.User{}
}

//注册用户
func (this *UserInfo) Register(id int) bool {
	this.userModel.CheckUserExist("aaa")

	password := utils.Md5("12345678")
	fmt.Println(password)

	return false
}
