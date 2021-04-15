package mysql

import (
	"database/sql"
	"time"

	"github.com/miaogu-go/bluebell/models"
)

const (
	TableNamePost = "post"
)

type Post struct {
	Id          int64  `json:"id" db:"id"`
	PostId      int64  `json:"post_id" db:"post_id"`
	AuthorId    int64  `json:"author_id" db:"author_id"`
	CommunityId int64  `json:"community_id" db:"community_id"`
	Title       string `json:"title" db:"title"`
	CreateTime  string `json:"create_time" db:"create_time"`
	UpdateTime  string `json:"update_time" db:"update_time"`
	Content     string `json:"content" db:"content"`
	Status      int    `json:"status" db:"status"`
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

// GetPostById 根据文章id获取文章
func GetPostById(postId int64) (*Post, error) {
	sqlStr := "SELECT * FROM " + TableNamePost + " WHERE post_id=?"
	post := new(Post)
	err := db.Get(post, sqlStr, postId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return post, nil
}
