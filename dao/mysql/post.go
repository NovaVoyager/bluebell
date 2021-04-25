package mysql

import (
	"database/sql"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

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

// GetPosts 获取帖子列表
func GetPosts(page, pageSize uint) ([]Post, error) {
	offset := (page - 1) * pageSize
	sqlStr := "SELECT * FROM " + TableNamePost + " ORDER BY create_time DESC LIMIT ?,?"
	posts := make([]Post, 0)
	err := db.Select(&posts, sqlStr, offset, pageSize)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostsByIds 根据文章ids获取文章列表
func GetPostsByIds(postIds []string) ([]Post, error) {
	if postIds == nil || len(postIds) == 0 {
		return nil, nil
	}
	sqlStr := "SELECT * FROM " + TableNamePost + " WHERE post_id IN (?) ORDER BY FIND_IN_SET(post_id,?)"
	query, args, err := sqlx.In(sqlStr, postIds, strings.Join(postIds, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	posts := make([]Post, 0)
	err = db.Select(&posts, query, args...)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
