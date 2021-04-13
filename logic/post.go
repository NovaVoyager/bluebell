package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(c *gin.Context, param *models.CreatePostReq) error {
	param.PostId = snowflake.GetID()
	err := mysql.CreatePost(param)
	if err != nil {
		return err
	}

	return nil
}
