package arrayproblems

import "fmt"

func SortDesc() {
	arr := []int{11, 3, 5, 6, 9, 8, 6, 4, 2}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("array in sorted order", arr)
}
