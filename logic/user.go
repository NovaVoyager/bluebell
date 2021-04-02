package logic

import (
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

func Signup(param *models.SignupReq) {
	mysql.QueryUserByUsername()
	snowflake.GetID()
	mysql.CreateUser()
}
