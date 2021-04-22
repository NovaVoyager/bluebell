package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/dao/redis"
	"github.com/miaogu-go/bluebell/models"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
)

type PostDetail struct {
	AuthorName string `json:"author_name"`
	mysql.Post
	Community mysql.Community `json:"community"`
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context, param *models.CreatePostReq) error {
	param.PostId = snowflake.GetID()
	err := mysql.CreatePost(param)
	if err != nil {
		return err
	}
	// 保存发布时间
	err = redis.SavePostPublishTime(param.PostId)
	if err != nil {
		return err
	}
	return nil
}

// GetPostDetail 获取帖子详情
func GetPostDetail(c *gin.Context, postId int64) (*PostDetail, error) {
	post, err := mysql.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	user, err := mysql.GetUserByUserId(post.AuthorId)
	if err != nil {
		return nil, err
	}
	community, err := mysql.GetCommunityDetailById(post.CommunityId)
	if err != nil {
		return nil, err
	}
	postDetail := &PostDetail{
		AuthorName: user.Username,
		Post:       *post,
		Community:  *community,
	}

	return postDetail, nil
}

// GetPosts 帖子列表
func GetPosts(c *gin.Context, param *models.PostsReq) ([]mysql.Post, error) {
	posts, err := mysql.GetPosts(param.Page, param.PageSize)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPosts2 帖子列表2
func GetPosts2(c *gin.Context, param *models.PostsReq) ([]mysql.Post, error) {
	start := (param.Page - 1) * param.PageSize
	postIds := make([]string, 0)
	if param.OrderType == models.PostOrderTypeTime {
		postIds = redis.GetPostIdsByTime(int64(start), int64(param.PageSize))
		if postIds == nil || len(postIds) == 0 {
			return nil, nil
		}
	} else {
		postIds = redis.GetPostIdsByScore(int64(start), int64(param.PageSize))
		if postIds == nil || len(postIds) == 0 {
			return nil, nil
		}
	}
	posts, err := mysql.GetPostsByIds(postIds)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
