package array_test

import "testing"

func TestArrayInit (t *testing.T) {
	var arr1 [3]int  = [3]int{1,3,5}
	var arr2 [2]int
	arr2[1] = 5
	arr3 := [2]int{1,2}
	arr4 := [...]int{3,5,6}
	var arr5 [2][2]int = [2][2]int{{1,2},{2,1}}
	t.Log(arr1[0],arr1[1])
	t.Log(arr2[0],arr2[1])
	t.Log(arr3[0],arr3[1])
	t.Log(arr4[0],arr4[1])
	t.Log(arr5[0][1],arr5[1][1])

}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1,3,4,6}
	for i := 0; i < len(arr); i ++ {
		t.Log(arr[i])
	}
	for idx,e := range arr{
		t.Log(idx, e)
	}
	for _,e := range arr{
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1,2,2,5,6,8,9}
	arr_sec1 := arr[:]
	t.Log(arr_sec1)
	arr_sec2 := arr[3:]
	t.Log(arr_sec2)
	arr_sec3 := arr[:3]
	t.Log(arr_sec3)
	arr_sec4 := arr[5:len(arr)]
	t.Log(arr_sec4)
}