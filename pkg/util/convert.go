package util

import "strconv"

func Int(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func UInt64(s string) uint64 {
	ui, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return ui
}

func Str(i uint64) string {
	return strconv.FormatUint(i, 10)
}
