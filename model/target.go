package model

import "time"

type Target struct {
	TargetId         int64  `json:"target_id" gorm:"primaryKey"`
	Uid              int64  `json:"uid"`
	Content          string `json:"content"`
	Status           int8   `json:"status"`
	Pic              string `json:"pic"`
	ExpertStartTime  int64  `json:"expert_start_time"`
	ExpertFinishTime int64  `json:"expert_finish_time"`
	ReviewTime       int64  `json:"review_time"`
	ReviewResult     string `json:"review_result"`
	ReviewStar       int8   `json:"review_star"`
	CreateTime       int64  `json:"create_time"`
	UpdateTime       int64  `json:"update_time"`
}

// 添加请求
type TargetAddReq struct {
	Uid              int64     `form:"uid"`
	Content          string    `form:"content"`
	Pic              string    `form:"pic"`
	ExpertStartTime  time.Time `form:"expert_start_time" time_format:"2006-01-02" time_utc:"1"`
	ExpertFinishTime time.Time `form:"expert_finish_time" time_format:"2006-01-02" time_utc:"1"`
	Status           int8      `form:"status"`
}

type TargetUpdateReq struct {
	TargetId         int64     `form:"target_id"`
	Content          string    `form:"content"`
	Pic              string    `form:"pic"`
	ExpertStartTime  time.Time `form:"expert_start_time" time_format:"2006-01-02" time_utc:"1"`
	ExpertFinishTime time.Time `form:"expert_finish_time" time_format:"2006-01-02" time_utc:"1"`
}

type TargetReviewReq struct {
	TargetId     int64  `form:"target_id"`
	ReviewResult string `form:"review_result"`
	ReviewStar   int8   `form:"review_star"`
}

type TargetListReq struct {
	TargetId         int64     `json:"target_id"`
	Content          string    `json:"content"`
	ExpertStartTime  time.Time `json:"expert_start_time" time_format:"2006-01-02" time_utc:"1"`
	ExpertFinishTime time.Time `json:"expert_finish_time" time_format:"2006-01-02" time_utc:"1"`
	Status           int8      `json:"status"`
	Page             int       `form:"page"`
	Size             int       `form:"size"`
}
