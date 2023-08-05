package benchmark

import (
	"RiceDouyin/dao"
	"testing"
	"time"
)

func BenchmarkCompare(b *testing.B) {
	b.ResetTimer()
	dao.NewVideoInstance().GetVideoSumByTime(time.Now())
}

// func BenchmarkOrderCompare(b *testing.B) {
// 	Sometime, _ := time.Parse(time.RFC3339, "2023-08-05 17:04:53")
// 	b.ResetTimer()
// 	dao.NewVideoInstance().GetVideoSumByTimeB(Sometime)
// }