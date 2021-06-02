package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1:1, 2:2, 3:3, 4:4}
	t.Log(m1[2],len(m1))
	m2 := map[int]int{}
	t.Log(m2[1],len(m2))
	m3 := make(map[int]int, 10)
	t.Log(m3[3],len(m3))
	// map类型不可以使用cap操作
	// t.Log(cap(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	if v,ok := m[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3's value is not existing.")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[int]int{1:4, 2:3, 3:2, 4:1}
	for key,value := range m{
		t.Logf("key is %d; value is %d.",key,value)
	}
}
