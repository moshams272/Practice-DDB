package main

import (
	"fmt"
)

func Ex2() {

	// Take inputs from user...
	fmt.Print("Enter array size: ")
	var array_size int
	fmt.Scan(&array_size)

	arr := make([]int, array_size)

	for i := 0; i < array_size; i++ {
		fmt.Printf("Enter the %dth element: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter Your target: ")
	var target int
	fmt.Scan(&target)

	// Loop through the array...
	for i, _ := range arr {

		if i == array_size-1 { // to handle out of bounds error...
			fmt.Println("The target not found...")
			break
		}

		// making sure that the target is the sum of two distinct values...
		if (arr[i]+arr[i+1] == target) && (arr[i] != arr[i+1]) {
			fmt.Printf("Target is found with the distinct indices i:%d, and j:%d", i, i+1)
			break
		}
	}

	/*
		Try to solve this problem to find all possible solutions in the array...
	*/

}
