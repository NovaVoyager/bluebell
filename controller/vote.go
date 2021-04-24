package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/miaogu-go/bluebell/logic"
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
	userId, err := GetUserInfo(c)
	if err != nil {
		zap.L().Error("GetUserInfo fail", zap.Error(err))
		ResponseError(c, CodeTokenInvalid)
		return
	}
	err = logic.VoteForPost(c, userId, param)
	if err != nil {
		zap.L().Error("logic.VoteForPost(c, userId) failed", zap.Error(err), zap.Int64("postId", param.PostId),
			zap.Int64("userId", userId), zap.Int8("direction", param.Direction))
		if errors.Is(err, logic.ErrVoteTimeExpire) {
			ResponseError(c, CodeVoteExpire)
			return
		}
		if errors.Is(err, logic.ErrVoted) {
			ResponseError(c, CodeNotRepeatVote)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
