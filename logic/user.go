package logic

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/miaogu-go/bluebell/pkg/jwt"

	"github.com/miaogu-go/bluebell/pkg/tools"

	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

const (
	PasswordSalt = "20210407160200"
)

var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorUserNotExist      = errors.New("用户不存在")
	ErrorInvalidPassword   = errors.New("密码错误")
	ErrorGenerateTokenFail = errors.New("token 生成失败")
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

type TokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login 登录
func Login(param *models.LoginReq) (*TokenResp, error) {
	user, err := mysql.QueryUserByUsername(param.User)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrorUserNotExist
	}
	if user.Password != tools.EncryptPassword(param.Password, PasswordSalt) {
		return nil, ErrorInvalidPassword
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.Username, user.UserId)
	if err != nil {
		return nil, ErrorGenerateTokenFail
	}
	loginResp := &TokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginResp, nil
}

// RefreshToken 刷新token
func RefreshToken(c *gin.Context, aToken, rToken string) (*TokenResp, error) {
	newAToken, newRToken, err := jwt.RefreshToken(aToken, rToken)
	if err != nil {
		return nil, err
	}
	tokenResp := &TokenResp{
		AccessToken:  newAToken,
		RefreshToken: newRToken,
	}

	return tokenResp, nil
}
