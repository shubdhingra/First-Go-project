package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 10)
	c2 := make(chan string, 10)

	go writer1("Sheep", c1)
	go writer2("Elephant", c2)
	go reader(c1, c2)

	fmt.Scanln()

}

func writer1(s1 string, c1 chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Adding element in c1")
		c1 <- s1
		time.Sleep(time.Millisecond * 100)
	}

}

func writer2(s2 string, c2 chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Adding element in c2")
		c2 <- s2
		time.Sleep(time.Millisecond * 20)
	}
}

func reader(c1 chan string, c2 chan string) {
	// for msg := range c1 {
	// 	fmt.Println("in reader ", msg)    // cannot be done because loop1 will always run and loop2 will never get executed due to selfish processing by loop1.
	// }

	// for msg := range c2 {
	// 	fmt.Println("in reader ", msg)
	// }

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
