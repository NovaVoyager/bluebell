package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	UserInfoKey = "user_id"
)

var ErrUserNotLogin = errors.New("用户未登录")

// GetUserInfo 获取用户登录信息
func GetUserInfo(c *gin.Context) (int64, error) {
	userId, exist := c.Get(UserInfoKey)
	if !exist {
		return 0, ErrUserNotLogin
	}
	uid, ok := userId.(int64)
	if !ok {
		return 0, ErrUserNotLogin
	}

	return uid, nil
}
