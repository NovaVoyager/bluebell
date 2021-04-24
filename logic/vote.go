package logic

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/dao/redis"
	"github.com/miaogu-go/bluebell/models"
)

const (
	OneWeekInSeconds = 7 * 86400
	ScorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoted          = errors.New("请勿重复投票")
)

// VoteForPost 投票
func VoteForPost(c *gin.Context, userId int64, param *models.VoteReq) error {
	postTime := redis.GetPostPublishTime(fmt.Sprintf("%d", param.PostId))
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	ov := redis.GetPostVoteUser(param.PostId, fmt.Sprintf("%d", userId))
	if ov == float64(param.Direction) {
		return ErrVoted
	}
	var dir float64
	if float64(param.Direction) > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - float64(param.Direction))
	score := dir * diff * ScorePerVote
	err := redis.SetPostScore(fmt.Sprintf("%d", param.PostId), score)
	if err != nil {
		return err
	}
	if param.Direction == 0 { //取消投票
		err = redis.RemUserVoteRecord(param.PostId, userId)
		if err != nil {
			return err
		}
	} else {
		err = redis.SaveUserVoteRecord(param.PostId, userId, float64(param.Direction))
		if err != nil {
			return err
		}
	}

	return nil
}
