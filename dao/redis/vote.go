package redis

import "github.com/go-redis/redis"

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
