package time_test

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 获取时间戳 int64
	timestamp := time.Now().Unix()
	t.Logf("unix timestamp: %d", timestamp)

	// 格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	tmstring := tm.Format("2006-01-02 03:04:05 PM")

	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	t.Log(tm.Format("2006-01-02 03:04:05 PM"))
	t.Log(tm.Format("02/01/2006 15:04:05 PM"))
	t.Log(time.Now().Format("2006/01/02/ 15:04:05"))
	t.Log(tmstring)

	// 从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	t.Log(tm2.Unix())

}
