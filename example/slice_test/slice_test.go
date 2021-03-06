/*
切片（slice）的使用
*/
package slice_test

import (
	"testing"
)

// 初始化切片
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) // 打印切片长度len、容量cap

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 3, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	s2 = append(s2, 1)
	t.Log(s2[1], s2[3])
}

// 切片长度自动扩容，前一次的2倍
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片共享内容
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unKonw"
	t.Log(Q2)
	t.Log(summer)
	t.Log(year)
}

// 切片比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 3, 5}

	// b := []int{1, 3, 5}
	// slice can only be compared to nil
	// if a == b {
	//	t.Log("equal")
	// }

	if a == nil {
		t.Log("equal")
	}
}
