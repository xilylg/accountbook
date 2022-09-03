package model

type User struct {
	Uid        int64  `json:"uid" gorm:"primaryKey"`
	Nickname   string `json:"nickname" gorm:"unique"`
	Password   string `json:"password"`
	Gender     int8   `json:"gender"`
	Birthday   string `json:"birthday"`
	Status     int8   `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type UserAddReq struct {
	Nickname   string `form:"nickname"`
	Password   string `form:"password"`
	Gender     int8   `form:"gender"`
	Birthday   string `form:"birthday"`
	CreateTime int64  `form:"create_time"`
}

type UserUpdateReq struct {
	Uid        int64  `form:"uid"`
	Nickname   string `form:"nickname"`
	Password   string `form:"password"`
	Gender     int8   `form:"gender"`
	Birthday   string `form:"birthday"`
	UpdateTime int64  `form:"update_time"`
}
