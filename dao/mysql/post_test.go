package mysql

import (
	"testing"

	"github.com/miaogu-go/bluebell/settings"

	"github.com/miaogu-go/bluebell/models"
)

func init() {
	conf := settings.DbConf{
		Host:        "127.0.0.1",
		Port:        3306,
		User:        "root",
		Password:    "123456",
		DbName:      "bluebell",
		MaxOpenConn: 50,
		MaxIdleConn: 10,
	}
	Init(&conf)
}

func TestCreatePost(t *testing.T) {
	post := &models.CreatePostReq{
		Title:       "test",
		Content:     "just a test",
		CommunityId: 1,
		AuthorId:    123,
		PostId:      100,
	}
	err := CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost(post) failed, err:%#v\n", err)
	}
	t.Log("create post success")
}
