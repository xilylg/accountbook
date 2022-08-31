package model

type AccountBook struct {
	BookId     int64  `db:"book_id"`
	Name       string `db:"name"`
	Cover      string `db:"cover"`
	Uid        int64  `db:"uid"`
	TargetId   int64  `db:"target_id"`
	CycleType  int8   `db:"cycle_type"`
	CycleStart string `db:"cycle_start"`
	CycleEnd   string `db:"cycle_end"`
	CreateTime int    `db:"create_time"`
	UpdateTime int    `db:"update_time"`
	Status     int8   `db:"status"`
}
