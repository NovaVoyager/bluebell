package logic

import (
	"errors"

	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

func Signup(param *models.SignupReq) error {
	userIsExist, err := mysql.CheckUserExist(param.User)
	if err != nil {
		return err
	}
	if userIsExist {
		return errors.New("用户已存在")
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
