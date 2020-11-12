package util

import "time"

const (
	TimeLayout = "2006-01-02 15:04:05"
)

func ToTimeString(i int64) string {
	t := time.Unix(i, 0)
	return t.Format(TimeLayout)
}
