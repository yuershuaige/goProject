package switch_condition

import "testing"

func TestSwitchMultiCase(t *testing.T) {
	// go中case可以支持多个值 并且每个case中自带break
	for i := 0; i < 5; i ++ {
		switch i {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not in 0 - 3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	// go中switch case可以构成 if else if 的效果
	for i := 0; i < 5; i ++ {
		switch  {
		case i % 2 == 0:
			t.Log("Even")
		case i % 2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}