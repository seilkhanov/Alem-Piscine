package piscine

func Swap(a *int, b *int) {
	var first int = *a
	var second int = *b
	*a = second
	*b = first
}
