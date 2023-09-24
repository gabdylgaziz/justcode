package main

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := merge(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}
