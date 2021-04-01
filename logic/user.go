package logic

import (
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

func Signup() {
	mysql.QueryUserByUsername()
	snowflake.GetID()
	mysql.CreateUser()
}
