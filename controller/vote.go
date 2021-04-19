package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/miaogu-go/bluebell/models"
	"go.uber.org/zap"
)

// VoteHandler 投票
func VoteHandler(c *gin.Context) {
	param := new(models.VoteReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	/*err := logic.Signup(param)
	if err != nil {
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		zap.L().Error("注册失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}*/
	ResponseSuccess(c, nil)
}
