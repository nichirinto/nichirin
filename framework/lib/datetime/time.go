package datetime

import "time"

func Unix() int64 {
	return time.Now().UnixMilli()
}

func UnixToFormat(v int64) string {
	return time.UnixMilli(v).Format(`2006-01-02 15:04:05`)
}
