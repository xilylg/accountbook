package model

type Book struct {
	BookId     int64  `json:"book_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Cover      string `json:"cover"`
	Uid        int64  `json:"uid"`
	TargetId   int64  `json:"target_id"`
	CycleType  int8   `json:"cycle_type"`
	CycleStart string `json:"cycle_start"`
	CycleEnd   string `json:"cycle_end"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Status     int8   `json:"status"`
}

// 添加请求
type BookAddReq struct {
	Name       string `form:"name" binding:"required"`
	Cover      string `form:"cover"`
	Uid        int64  `form:"uid" binding:"numeric"`
	TargetId   int64  `form:"target_id" binding:"numeric"`
	CycleType  int8   `form:"cycle_type" binding:"numeric"`
	CycleStart string `form:"cycle_start"`
	CycleEnd   string `form:"cycle_end"`
}

type BookUpdateReq struct {
	BookId     int64  `form:"book_id"`
	Name       string `form:"name"`
	Cover      string `form:"cover"`
	TargetId   int64  `form:"target_id"`
	CycleType  int8   `form:"cycle_type"`
	CycleStart string `form:"cycle_start"`
	CycleEnd   string `form:"cycle_end"`
}

type BookListReq struct {
	BookId int64  `form:"book_id"`
	Name   string `form:"name"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}
