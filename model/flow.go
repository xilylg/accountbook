package model

type AccountFlow2022 struct {
	FlowId     int64  `db:"flow_id"`
	Cost       int    `db:"cost"`
	Uid        int64  `db:"uid"`
	BookId     int64  `db:"book_id"`
	MemberId   int64  `db:"member_id"`
	Tip        string `db:"tip"`
	ReviewStar int8   `db:"review_star"`
	ReviewTime int    `db:"review_time"`
	CreateTime int    `db:"create_time"`
}
