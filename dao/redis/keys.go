package redis

const (
	KeyPrefix    = "bluebell:"
	KeyPostTime  = "post:time"     //zset
	KeyPostScore = "post:score"    //zset
	KeyPostVoted = "post:voted:%d" //zset; 参数为 postId
)
