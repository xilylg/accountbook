package utils

import "time"

func Day2Unix(str string) int64 {
	return Str2Unix(str, "2006-01-02")
}

func Time2Unix(str string) int64 {
	return Str2Unix(str, "2006-01-02 15:04:05")
}

// 2006-01-02 15:04:05
func Str2Unix(str string, layout string) int64 {
	loc, _ := time.LoadLocation("Local")
	unix, _ := time.ParseInLocation(layout, str, loc)
	return unix.Unix()
}
