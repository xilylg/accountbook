package model

// 数据表
type Member struct {
	MemberId   int64  `gorm:"primaryKey" json:"member_id"`
	Uid        int64  `json:"uid"`
	Nickname   string `json:"nickname"`
	CreateTime int    `json:"create_time"`
	Relation   string `json:"relation"`
	Status     int8   `json:"status"`
}

// 添加请求
type MemberAddReq struct {
	Uid      int64  `form:"uid" binding:"required,"`
	Nickname string `form:"nickname" binding:"required,"`
	Relation string `form:"relation"`
}
