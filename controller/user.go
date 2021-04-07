package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"github.com/miaogu-go/bluebell/models"
)

func SignUpHandler(c *gin.Context) {
	param := new(models.SignupReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": removeTopStruct(errs.Translate(trans))})
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
