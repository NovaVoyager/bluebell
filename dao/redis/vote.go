package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// GetPostPublishTime 获取文章发布时间
func GetPostPublishTime(postId string) float64 {
	postTime := rdb.ZScore(GetKeyPostTime(), postId).Val()

	return postTime
}

// GetPostVoteUser 获取文章投票记录
func GetPostVoteUser(postId int64, userId string) float64 {
	oldValue := rdb.ZScore(GetKeyPostVoted(postId), userId).Val()

	return oldValue
}

// SetPostScore 设置文章分数
func SetPostScore(postId string, score float64) error {
	err := rdb.ZIncrBy(GetKeyPostScore(), score, postId).Err()
	if err != nil {
		return err
	}

	return nil
}

// SaveUserVoteRecord 保存用户投票记录
func SaveUserVoteRecord(postId, userId int64, direction float64) error {
	z := redis.Z{
		Score:  direction,
		Member: userId,
	}
	err := rdb.ZAdd(GetKeyPostVoted(postId), z).Err()
	if err != nil {
		return err
	}

	return nil
}

// RemUserVoteRecord 移除用户投票记录
func RemUserVoteRecord(postId, userId int64) error {
	err := rdb.ZRem(GetKeyPostVoted(postId), userId).Err()
	if err != nil {
		return err
	}

	return nil
}

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
