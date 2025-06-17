package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("array in sorted order", arr)
}
