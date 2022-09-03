package model

type Item struct {
	ItemId     int64  `json:"item_id" gorm:"primaryKey"`
	Uid        int64  `json:"uid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Flow       int8   `json:"flow"`
	Status     int8   `json:"status"`
}

// 添加请求
type ItemAddReq struct {
	Uid  int64  `form:"uid"`
	Name string `form:"name"`
	Flow int8   `form:"flow"`
}

type ItemUpdateReq struct {
	ItemId int64  `form:"item_id"`
	Uid    int64  `form:"uid"`
	Name   string `form:"name"`
	Flow   int8   `form:"flow"`
}

type ItemListReq struct {
	ItemId int64  `form:"Item_id"`
	Name   string `form:"name"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}
