package evenodd

func Even(v int) bool {
	return v%2 == 0
}

func Odd(v int) bool {
	return !Even(v)
}
