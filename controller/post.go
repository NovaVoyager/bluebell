package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/miaogu-go/bluebell/logic"
	"github.com/miaogu-go/bluebell/models"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	param := new(models.CreatePostReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("CreatePostHandler failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
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
	param.AuthorId = userId
	err = logic.CreatePost(c, param)
	if err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
