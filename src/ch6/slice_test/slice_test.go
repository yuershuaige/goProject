package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var slice1 []int
	t.Log(len(slice1),cap(slice1))
	slice1 = append(slice1, 1,2)
	t.Log(len(slice1),cap(slice1))

	var slice2 []int = []int{1,2,3,4}
	t.Log(len(slice2),cap(slice2))
	slice2 = append(slice2,1,1,1)
	t.Log(len(slice2),cap(slice2))
	slice2 = append(slice2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1)
	t.Log(len(slice2),cap(slice2))

	slice3 := make([]int,3,5)
	t.Log(len(slice3),cap(slice3))
	//只初始化len个元素为0  capacity指预留的总容量  len指已经填充的元素个数
	//t.Log(slice3[0],slice3[1],slice3[2],slice[3])
	t.Log(slice3[0],slice3[1],slice3[2])
	slice3 = append(slice3,3)
	t.Log(len(slice3),cap(slice3))
	t.Log(slice3[0],slice3[1],slice3[2],slice3[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	// 每次slice扩容都是 2的n（0，1，2，3）次方大小
	for i := 0; i <= 10; i++ {
		// 因为其扩容所以 s 指向的连续存储空间会发生变化 所以需要把新空间地址重新赋值给s
		s = append(s,1)
		t.Log(len(s),cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan","Feb","Mar","Apr","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	s1 := year[3:6]
	t.Log(s1, len(s1), cap(s1))
	s2 := year[5:8]
	t.Log(s2, len(s2), cap(s2))
	s2[0] = "unknown"
	t.Log(s1)
	t.Log(s2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	s1 := []int{1,2,3,4}
	s2 := []int{1,2,3,4}
	t.Log(s1,s2)
	// slice和数组不一样 不支持同维数下比较元素 只可以和nil进行比较
	//if (s1 == s2) {
	//	t.Log("equals is true")
	//}
}
