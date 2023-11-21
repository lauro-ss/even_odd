package evenodd

import "testing"

func TestOdd(t *testing.T) {
	v := 2
	if Odd(v) {
		t.Errorf("The value %v is even", v)
	}
}

func TestEven(t *testing.T) {
	v := 3
	if Even(v) {
		t.Errorf("The value %v is odd", v)
	}
}
