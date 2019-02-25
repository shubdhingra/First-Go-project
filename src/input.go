package main

import (
	"fmt"
)

func main() {
	var ans int
	fmt.Println("Enter a String")
	status, err := fmt.Scanln(&ans)
	if err != nil {
		fmt.Printf(err)
		fmt.Printf("%d", status)
	} else {
		fmt.Printf("%d", ans)
	}

}
