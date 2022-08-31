package model

type User struct {
	Uid        uint64 `gorm:"primaryKey"`
	Nickname   string `gorm:"unique"`
	Password   string `db:"password"`
	Gender     int8   `db:"gender"`
	Birthday   string `db:"birthday"`
	Status     int8   `db:"status"`
	CreateTime int    `db:"create_time"`
	UpdateTime int    `db:"update_time"`
}
