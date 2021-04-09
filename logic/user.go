package logic

import (
	"errors"

	"github.com/miaogu-go/bluebell/pkg/tools"

	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

const (
	PasswordSalt = "20210407160200"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// Signup 注册
func Signup(param *models.SignupReq) error {
	userIsExist, err := mysql.CheckUserExist(param.User)
	if err != nil {
		return err
	}
	if userIsExist {
		return ErrorUserExist
	}
	userId := snowflake.GetID()
	u := &models.User{
		UserId:   userId,
		Username: param.User,
		Password: param.Password,
	}
	err = mysql.CreateUser(u)
	if err != nil {
		return err
	}

	return nil
}

// Login 登录
func Login(param *models.LoginReq) error {
	user, err := mysql.QueryUserByUsername(param.User)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrorUserNotExist
	}
	if user.Password != tools.EncryptPassword(param.Password, PasswordSalt) {
		return ErrorInvalidPassword
	}

	return nil
}
