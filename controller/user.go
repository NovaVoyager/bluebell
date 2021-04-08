package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"github.com/miaogu-go/bluebell/models"
)

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	param := new(models.SignupReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if ok {
			c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(errs.Translate(trans))})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	err := logic.Signup(param)
	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	//验证参数
	param := new(models.LoginReq)
	if err := c.ShouldBind(param); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(errs.Translate(trans))})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	err := logic.Login(param)
	if err != nil {
		zap.L().Error("LoginHandler failed", zap.String("username", param.User), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": "用户名或密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
}
