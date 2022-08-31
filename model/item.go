package model

type AccountItem struct {
	ItemId     int64  `db:"item_id"`
	Uid        int64  `db:"uid"`
	ItemName   string `db:"item_name"`
	CreateTime int    `db:"create_time"`
	Flow       int8   `db:"flow"`
}
