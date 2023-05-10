package time

import (
	"time"
)

// GetNowTime 暂时不用string格式
func GetNowTime() string {
	now := time.Now()
	//格式化时间输出
	nowTime := now.Format("2006-01-02 15:04:05")
	return nowTime
}
