package mysql

import "database/sql"

const (
	TableNameCommunity = "community"
)

type Community struct {
	Id            uint32 `json:"id" db:"id"`
	CommunityId   uint32 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction  string `json:"introduction" db:"introduction"`
	CreateTime    string `json:"create_time" db:"create_time"`
	UpdateTime    string `json:"update_time" db:"update_time"`
}

// GetCommunityList 获取社区列表
func GetCommunityList() ([]Community, error) {
	sqlStr := "SELECT * FROM " + TableNameCommunity
	communities := make([]Community, 0)
	err := db.Select(&communities, sqlStr)
	if err != nil {
		return nil, err
	}

	return communities, nil
}

// GetCommunityDetailById 根据社区id获取详情
func GetCommunityDetailById(id int64) (*Community, error) {
	sqlStr := "SELECT * FROM " + TableNameCommunity + " WHERE community_id=?"
	community := new(Community)
	err := db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return community, nil
}
