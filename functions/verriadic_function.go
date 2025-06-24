package functions

func VeriadicFunction(a ...int) int {
	var sum int
	for _, num := range a {
		sum += num
	}
	return sum
}
