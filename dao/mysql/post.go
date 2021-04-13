package mysql

import (
	"time"

	"github.com/miaogu-go/bluebell/models"
)

const (
	TableNamePost = "post"
)

type Post struct {
	Id          int64  `json:"id" ddb:""`
	PostId      int64  `json:"post_id" ddb:"post_id"`
	AuthorId    int64  `json:"author_id" ddb:"author_id"`
	CommunityId int64  `json:"community_id" ddb:"community_id"`
	Title       string `json:"title" ddb:"title"`
	CreateTime  string `json:"create_time" ddb:"create_time"`
	UpdateTime  string `json:"update_time" ddb:"update_time"`
	Content     string `json:"content" ddb:"content"`
	Status      int    `json:"status" ddb:"status"`
}

// CreatePost 创建帖子
func CreatePost(param *models.CreatePostReq) error {
	sqlStr := "INSERT INTO " + TableNamePost + " (`post_id`,`title`,`content`,`author_id`,`community_id`,`create_time`," +
		"`update_time`)VALUES(?,?,?,?,?,?,?)"
	date := time.Now()
	_, err := db.Exec(sqlStr, param.PostId, param.Title, param.Content, param.AuthorId, param.CommunityId, date, date)
	if err != nil {
		return err
	}

	return nil
}
