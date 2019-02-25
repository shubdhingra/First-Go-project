package main

import (
	"fmt"
)

func main() {
	count := 8
	var arr [15]int
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < count; i++ {

		arr[i] = arr[i-1] + arr[i-2]
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d", arr[i])
	}
}
