package model

type AccountTarget struct {
	TargetId         int64  `db:"target_id"`
	Uid              int64  `db:"uid"`
	Content          string `db:"content"`
	Status           string `db:"status"`
	Pic              string `db:"pic"`
	ExpertStartTime  int    `db:"expert_start_time"`
	ExpertFinishTime int    `db:"expert_finish_time"`
	ReviewTime       int    `db:"review_time"`
	ReviewResult     string `db:"review_result"`
	ReviewStar       int8   `db:"review_star"`
	CreateTime       int    `db:"create_time"`
	UpdateTime       int    `db:"update_time"`
}
