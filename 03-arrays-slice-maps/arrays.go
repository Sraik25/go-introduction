package main

import "fmt"

func arraysImplement() {
	var arr [5]int
	fmt.Println("empty: ", arr)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("initializated array: ", b)

	arr[2] = 200

	fmt.Println("arr: ", arr)
	fmt.Println("arr[2]: ", arr[2])
}
