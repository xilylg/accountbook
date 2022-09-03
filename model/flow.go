package model

type Flow struct {
	FlowId     int64  `json:"flow_id" gorm:"primaryKey"`
	Cost       int    `json:"cost"`
	Uid        int64  `json:"uid"`
	BookId     int64  `json:"book_id"`
	MemberId   int64  `json:"member_id"`
	Tip        string `json:"tip"`
	ReviewStar int8   `json:"review_star"`
	ReviewTime int64  `json:"review_time"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Status     int8   `json:"status"`
}

// 添加请求
type FlowAddReq struct {
	Cost       int    `form:"cost"`
	Uid        int64  `form:"uid"`
	BookId     int64  `form:"book_id"`
	MemberId   int64  `form:"member_id"`
	Tip        string `form:"tip"`
	CreateTime int64  `form:"create_time"`
}

type FlowUpdateReq struct {
	FlowId     int64  `form:"flow_id"`
	Cost       int    `form:"cost"`
	BookId     int64  `form:"book_id"`
	MemberId   int64  `form:"member_id"`
	Tip        string `form:"tip"`
	UpdateTime int64  `form:"update_time"`
}

type FlowReviewReq struct {
	FlowId     int64  `form:"flow_id"`
	Tip        string `form:"tip"`
	ReviewStar int8   `form:"review_star"`
	ReviewTime int64  `form:"review_time"`
	UpdateTime int64  `form:"update_time"`
}

type FlowListReq struct {
	FlowId   int64 `form:"flow_id"`
	BookId   int64 `form:"book_id"`
	MemberId int64 `form:"member_id"`
	Status   int8  `form:"status"`
	Page     int   `form:"page"`
	Size     int   `form:"size"`
}
