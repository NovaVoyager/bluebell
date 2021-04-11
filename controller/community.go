package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"go.uber.org/zap"
)

// CommunityHandler 获取社区列表
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList(c)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
