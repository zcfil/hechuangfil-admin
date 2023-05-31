package common

import (
	"time"
)


func TimeToDay(time time.Time) int32 {
	year := time.Year()
	month := time.Month()
	day := time.Day()
	return int32(year) * 10000 + int32(month) * 100 + int32(day)
}