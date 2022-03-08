package ch6

import "testing"

func TestArraySyntax(t *testing.T) {
	var a [3]int
	a[0] = 1

	b := [...]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}

	t.Log(a, b, c)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{2, 3, 4, 5, 6, 7}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArrayLength(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(len(a))
	// 针对数组的截取相当于是：左包含，右不包含
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:3])
}

// 切片数据结构
func TestSliceInit(t *testing.T) {
	// 是一个结构体，一个指针，第二个是元素的个数，第三个 内部数组的容量
	// 可变长，不声明长度
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[3])
}

// 切片如何实现可变长的
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片共享存储结构,一个修改另一个也会跟着修改
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dev"}
	Q2 := year[3:6]
	t.Log(Q2)
	summer := year[5:8]
	t.Log(summer)
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(summer)

	summer1 := year[5:8]
	// slice can only be compared to nil
	//if summer == summer1 {
	if summer1 == nil {
		t.Log("summer1 is nil")
	}
}
