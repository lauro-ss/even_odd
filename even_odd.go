package evenodd

func even(v int) bool {
	return v%2 == 0
}

func odd(v int) bool {
	return !even(v)
}
