package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go simple("first")
	go simple("second")
	fmt.Scanln()

}

func simple(str string) {
	for i := 0; i < 20; i++ {
		fmt.Println(str + strconv.Itoa(i))
		time.Sleep(2000)
	}
}
