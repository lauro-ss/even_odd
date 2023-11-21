package evenodd

import "testing"

func TestOdd(t *testing.T) {
	v := 2
	if odd(v) {
		t.Errorf("The value %v is even", v)
	}
}

func TestEven(t *testing.T) {
	v := 3
	if even(v) {
		t.Errorf("The value %v is odd", v)
	}
}
