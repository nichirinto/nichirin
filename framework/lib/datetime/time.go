package datetime

import "time"

func Unix() int64 {
	return time.Now().UnixMilli()
}
