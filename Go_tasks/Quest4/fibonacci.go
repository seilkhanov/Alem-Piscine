package piscine

func Fibonacci(index int) int {
	a := 0
	b := 1
	result := 0

	if index < 0 {
		return -1
	}

	if index == 0 || index == 1 {
		return 1
	}

	for i := 2; i <= index; i++ {
		result = a + b
		a = b
		b = result
	}
	return result
}
