package arrayproblems

import (
	"fmt"
	"math"
)

func SecondLargest() {
	arr := []int{10, 10, 10}
	max1 := math.MinInt
	max2 := math.MinInt
	for _, num := range arr {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 && num != max1 {
			max2 = num
		}
	}
	fmt.Println("Second largest number", max2)

}
