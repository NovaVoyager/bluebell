package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/dao/mysql"
)

// GetCommunityList 获取社区列表
func GetCommunityList(c *gin.Context) ([]mysql.Community, error) {
	communities, err := mysql.GetCommunityList()
	if err != nil {
		return nil, err
	}

	return communities, nil
}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(c *gin.Context, id int64) (*mysql.Community, error) {
	community, err := mysql.GetCommunityDetailById(id)
	if err != nil {
		return nil, err
	}

	return community, nil
}
