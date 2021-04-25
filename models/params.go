package models

type SignupReq struct {
	User       string `json:"user" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginReq struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreatePostReq struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CommunityId int64  `json:"community_id" binding:"required"`
	AuthorId    int64  `json:"author_id"`
	PostId      int64  `json:"post_id"`
}

type PostsReq struct {
	Page        uint  `json:"page"`
	PageSize    uint  `json:"page_size"`
	OrderType   int8  `json:"order_type"` // 1 时间，2 点赞
	CommunityId int64 `json:"community_id"`
}

type VoteReq struct {
	PostId    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction" binding:"oneof=1 0 -1"`
}
