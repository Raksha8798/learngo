package test

import "testing"

func TestAdd(t *testing.T) {
	input := Add(3, 4)
	output := 7
	if input != output {
		t.Errorf("not perfect")
	}
}
func TestSub(t *testing.T) {
	input := Sub(4, 3)
	output := 1
	if input != output {
		t.Errorf("not perfect")
	}
}
func TestMulti(t *testing.T) {
	input := Multi(3, 4)
	output := 12
	if input != output {
		t.Errorf("not perfect")
	}
}
func TestDiv(t *testing.T) {
	input := Div(18, 6)
	output := 3
	if input != output {
		t.Errorf("not perfect")
	}
}
