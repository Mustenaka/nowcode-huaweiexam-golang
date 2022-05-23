package main

import (
	"fmt"
)

func main() {
	var N int = 0
	fmt.Scanln(&N)

	// Construct an array of length 500, with each element initialized to 0.
	array := make([]int, 500)
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scanln(&tmp)
		array[tmp]++
	}

	for i := 0; i < 500; i++ {
		if array[i] > 0 {
			fmt.Println(i)
		}
	}

}
