package redis

import (
	"fmt"
	"time"

	"github.com/miaogu-go/bluebell/models"

	"github.com/go-redis/redis"
)

// SavePostPublishTime 保存帖子发布时间
func SavePostPublishTime(postId int64) error {
	z := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	}
	err := rdb.ZAdd(GetKeyPostTime(), z).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetPostIdsByTime 根据时间获取帖子指定范围的ids
func GetPostIdsByTime(start, stop int64) []string {
	key := GetKeyPostTime()
	return rdb.ZRevRange(key, start, stop).Val()
}

// GetPostIdsByScore 根据分数获取帖子指定范围的ids
func GetPostIdsByScore(start, stop int64) []string {
	key := GetKeyPostScore()
	return rdb.ZRevRange(key, start, stop).Val()
}

// SaveCommunityPost 保存社区下的帖子数据
func SaveCommunityPost(communityId string, postId int64) error {
	key := GetKeyCommunityPost(communityId)
	return rdb.SAdd(key, postId).Err()
}

// GetCommunityPostIds 根据社区筛选postId
func GetCommunityPostIds(communityKey string, orderType int8, page, pageSize int64) ([]string, error) {
	orderKey := ""
	if orderType == models.PostOrderTypeTime { //发布时间
		orderKey = GetKeyPostTime()
	} else { //分数
		orderKey = GetKeyPostScore()
	}
	key := orderKey + fmt.Sprintf("%d", time.Now().Unix())
	pipe := rdb.Pipeline()
	pipe.ZInterStore(key, redis.ZStore{
		Aggregate: "MAX",
	}, communityKey, orderKey)
	pipe.Expire(key, time.Second*60)
	_, err := pipe.Exec()
	if err != nil {
		return nil, err
	}
	postIds := rdb.ZRevRange(key, page, pageSize).Val()

	return postIds, nil
}
