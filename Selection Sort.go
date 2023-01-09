package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 90, 20, 2, 4, 5}
	var min_index int

	for i := 0; i < len(arr); i++ {
		min_index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[i], arr[min_index] = arr[min_index], arr[i]

	}

	fmt.Println(arr)
}
