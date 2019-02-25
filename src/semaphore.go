package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	s := "aa"
	for i := 0; i < 11; i++ {
		fmt.Println(s + strconv.Itoa(i))
		time.Sleep(10000)
	}

	var wg sync.WaitGroup //just like a semaphore or lock being given
	wg.Add(1)
	go func() {
		wait1("first")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		wait1("second")
		wg.Done()
	}()
	wg.Wait()

}

func wait1(s string) {
	for i := 0; i < 11; i++ {
		fmt.Println(s + strconv.Itoa(i))
		time.Sleep(1000)
	}
}
