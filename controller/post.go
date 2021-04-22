package controller

import (
	"strconv"

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

// GetPostDetailHandler 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	postId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("Param id failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(c, postId)
	if err != nil {
		zap.L().Error("GetPostDetail(postId) failed", zap.Error(err), zap.Int64("postId", postId))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostsHandler 帖子列表
func GetPostsHandler(c *gin.Context) {
	param := new(models.PostsReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("GetPostsHandler failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPosts(c, param)
	if err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPosts2Handler 帖子列表2
func GetPosts2Handler(c *gin.Context) {
	param := new(models.PostsReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("GetPostsHandler failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPosts2(c, param)
	if err != nil {
		zap.L().Error("logic.GetPosts2(c, param) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
