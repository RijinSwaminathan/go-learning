// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	result := addNum(10, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
func addNum(a ...int) int {
	var sum int
	for _, num := range a {
		sum += num
	}
	return sum
}
