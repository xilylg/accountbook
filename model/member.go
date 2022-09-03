package model

// 数据表
type Member struct {
	MemberId   int64  `gorm:"primaryKey" json:"member_id" binding:"required,numeric"`
	Uid        int64  `json:"uid" binding:"required,numeric"`
	Nickname   string `json:"nickname"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Relation   string `json:"relation"`
	Status     int8   `json:"status"`
}

// 添加请求
type MemberAddReq struct {
	Uid      int64  `form:"uid"`
	Nickname string `form:"nickname" binding:"required"`
	Relation string `form:"relation"`
}

type MemberUpdateReq struct {
	MemberId int64  `form:"mid" binding:"numeric"`
	Nickname string `form:"nickname"`
	Relation string `form:"relation"`
}

type MemberListReq struct {
	MemberId int64  `form:"mid" binding:"numeric"`
	Nickname string `form:"nickname"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}
