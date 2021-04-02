package controller

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"github.com/miaogu-go/bluebell/models"
)

func SignUpHandler(c *gin.Context) {
	param := new(models.SignupReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": "参数错误"})
		return
	}
	logic.Signup(param)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
