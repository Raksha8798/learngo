package pack

import "testing"

func TestAdd(t *testing.T) {
	input := Add(12, 3)
	output := 15
	if input != output {
		t.Errorf("Not perfect")
	}
}
func TestStr(t *testing.T) {
	input := Str("hello navi")
	output := "hello navi"
	if input != output {
		t.Errorf("Not Perfect")
	}
}
