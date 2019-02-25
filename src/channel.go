package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 10)
	go writer("Sheep", c)
	go reader(c)

	fmt.Scanln()

}

func reader(c chan string) {
	for msg := range c { // to read the values published by the publisher dynamically
		fmt.Println("in reader ", msg)
	}
}

func writer(s string, c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("in writer ", i)
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
}
