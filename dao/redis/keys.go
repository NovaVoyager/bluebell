package redis

import "fmt"

const (
	KeyPrefix        = "bluebell:"
	KeyPostTime      = "post:time"         //zset; postId:publishTime,eg: 12347:11889901
	KeyPostScore     = "post:score"        //zset; postId:voteNum,eg:12347:1000
	KeyPostVoted     = "post:voted:%s"     //zset; 参数为 postId
	KeyCommunityPost = "post:community:%s" //set 参数为社区id
)

// GetKeyPostTime 获取文章发布时间redis key
func GetKeyPostTime() string {
	return KeyPrefix + KeyPostTime
}

// GetKeyPostScore 获取文章得票数redis key
func GetKeyPostScore() string {
	return KeyPrefix + KeyPostScore
}

// GetKeyPostVoted 获取文章下用户投票记录
func GetKeyPostVoted(postId string) string {
	return KeyPrefix + fmt.Sprintf(KeyPostVoted, postId)
}

// GetKeyPostVoted 获取文章下用户投票记录
func GetKeyCommunityPost(communityId string) string {
	return KeyPrefix + fmt.Sprintf(KeyCommunityPost, communityId)
}
